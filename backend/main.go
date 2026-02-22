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

	paymentRepo := payrepo.New(db)
	payUC := payuc.New(paymentRepo)

	apiHandler := &api.APIHandler{
		Auth: authH,
		Payment: payUC,
	}

	server := srv.NewServer(apiHandler, config.OpenapiYamlLocation)

	addr := config.HttpAddress
	log.Printf("starting server on %s", addr)
	server.Start(addr)
}

func initDB(db *sql.DB) error {
	stmts := []string{

		// USERS TABLE
		`CREATE TABLE IF NOT EXISTS users (
		  id INTEGER PRIMARY KEY AUTOINCREMENT,
		  email TEXT NOT NULL UNIQUE,
		  password_hash TEXT NOT NULL,
		  role TEXT NOT NULL
		);`,

		// PAYMENTS TABLE  ← TAMBAH INI
		`CREATE TABLE IF NOT EXISTS payments (
		  id TEXT PRIMARY KEY,
		  amount INTEGER NOT NULL,
		  status TEXT NOT NULL,
		  user_id TEXT NOT NULL
		);`,
	}

	for _, s := range stmts {
		if _, err := db.Exec(s); err != nil {
			return err
		}
	}

	// ===== SEED USERS =====
	var cnt int
	row := db.QueryRow("SELECT COUNT(1) FROM users")
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

	// ===== SEED PAYMENTS =====
	if _, err := db.Exec(`
	INSERT OR IGNORE INTO payments(id,amount,status,user_id) VALUES
	('pay_001',100000,'pending','1'),
	('pay_002',250000,'success','1');
	`); err != nil {
		return err
	}

	db.SetConnMaxLifetime(time.Minute * 5)
	return nil
}
