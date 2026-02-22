package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/durianpay/fullstack-boilerplate/internal/api"
	"github.com/durianpay/fullstack-boilerplate/internal/config"
	ah "github.com/durianpay/fullstack-boilerplate/internal/module/auth/handler"
	ar "github.com/durianpay/fullstack-boilerplate/internal/module/auth/repository"
	au "github.com/durianpay/fullstack-boilerplate/internal/module/auth/usecase"
	srv "github.com/durianpay/fullstack-boilerplate/internal/service/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	payuc "github.com/durianpay/fullstack-boilerplate/internal/module/payment/usecase"
	payrepo "github.com/durianpay/fullstack-boilerplate/internal/module/payment/repository"
)

func main() {
	_ = godotenv.Load()

	db, err := sql.Open("sqlite3", "dashboard.db?_foreign_keys=1")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := initDB(db); err != nil {
		log.Fatal(err)
	}

	JwtExpiredDuration, err := time.ParseDuration(config.JwtExpired)
	if err != nil {
		panic(err)
	}

	userRepo := ar.NewUserRepo(db)
	authUC := au.NewAuthUsecase(userRepo, config.JwtSecret, JwtExpiredDuration)
	authH := ah.NewAuthHandler(authUC)

	// INIT PAYMENT REPOSITORY DENGAN DB
	paymentRepo := payrepo.New(db)
	payUC := payuc.New(paymentRepo)

	apiHandler := &api.APIHandler{
		Auth:    authH,
		Payment: payUC,
	}

	server := srv.NewServer(apiHandler, config.OpenapiYamlLocation)
	addr := config.HttpAddress
	log.Printf("starting server on %s", addr)
	server.Start(addr)
}

func initDB(db *sql.DB) error {
	// Create users table
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL,
		role TEXT NOT NULL
	);`); err != nil {
		return err
	}

	// Create/update payments table with ALL columns
	// SQLite doesn't support ALTER TABLE ADD COLUMN easily for all cases,
	// so we recreate the table if it doesn't have the right schema
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS payments (
		id TEXT PRIMARY KEY,
		merchant TEXT NOT NULL,
		amount INTEGER NOT NULL,
		status TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		user_id TEXT NOT NULL
	);`); err != nil {
		return err
	}

	// MIGRATION: Add missing columns if table exists with old schema
	// Check if 'merchant' column exists
	// var hasMerchant int
	row := db.QueryRow(`PRAGMA table_info(payments);`)
	// Simple approach: try to query merchant, if fails, recreate table
	var testMerchant string
	err := db.QueryRow(`SELECT merchant FROM payments LIMIT 1`).Scan(&testMerchant)
	if err != nil && err != sql.ErrNoRows {
		// Column doesn't exist, recreate table
		db.Exec(`DROP TABLE IF EXISTS payments`)
		db.Exec(`CREATE TABLE payments (
			id TEXT PRIMARY KEY,
			merchant TEXT NOT NULL,
			amount INTEGER NOT NULL,
			status TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			user_id TEXT NOT NULL
		);`)
	}

	// Seed users
	var cnt int
	row = db.QueryRow("SELECT COUNT(1) FROM users")
	if err := row.Scan(&cnt); err != nil {
		return err
	}

	if cnt == 0 {
		hash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		db.Exec("INSERT INTO users(email,password_hash,role) VALUES (?,?,?)",
			"cs@test.com", string(hash), "cs")
		db.Exec("INSERT INTO users(email,password_hash,role) VALUES (?,?,?)",
			"operation@test.com", string(hash), "operation")
	}

	// SEED PAYMENTS
	if _, err := db.Exec(`
	INSERT OR IGNORE INTO payments(id,merchant,amount,status,created_at,user_id) VALUES
	('pay_001','Merchant A',100000,'processing','2026-02-20 10:00:00','1'),
	('pay_002','Merchant B',250000,'completed','2026-02-20 11:00:00','1'),
	('pay_003','Merchant C',75000,'failed','2026-02-20 12:00:00','1'),
	('pay_004','Merchant D',500000,'processing','2026-02-20 13:00:00','2'),
	('pay_005','Merchant E',150000,'completed','2026-02-20 14:00:00','2');
	`); err != nil {
		return err
	}

	db.SetConnMaxLifetime(time.Minute * 5)
	return nil
}