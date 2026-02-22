package repository

import (
	"database/sql"
	"fmt"
	"time"
)

type Payment struct {
	ID        string
	Merchant  string
	Amount    int64
	Status    string
	CreatedAt time.Time
	UserID    string
}

type Repo struct {
	DB *sql.DB
}

func New(db *sql.DB) *Repo {
	return &Repo{DB: db}
}

func (r *Repo) ListByUser(userID string) ([]Payment, error) {
	rows, err := r.DB.Query(
		`SELECT id, merchant, amount, status, created_at, user_id FROM payments WHERE user_id = ? ORDER BY created_at DESC`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []Payment
	for rows.Next() {
		var p Payment
		if err := rows.Scan(&p.ID, &p.Merchant, &p.Amount, &p.Status, &p.CreatedAt, &p.UserID); err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, rows.Err()
}

// Get payment by ID
func (r *Repo) GetByID(id string) (*Payment, error) {
	row := r.DB.QueryRow(
		`SELECT id, merchant, amount, status, created_at, user_id FROM payments WHERE id = ?`,
		id,
	)

	var p Payment
	if err := row.Scan(&p.ID, &p.Merchant, &p.Amount, &p.Status, &p.CreatedAt, &p.UserID); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("payment not found")
		}
		return nil, err
	}

	return &p, nil
}

// Update payment status
func (r *Repo) UpdateStatus(id string, status string) error {
	result, err := r.DB.Exec(
		`UPDATE payments SET status = ? WHERE id = ?`,
		status, id,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("payment not found")
	}

	return nil
}

// ✅ NEW: Filter by status
func (r *Repo) ListByUserAndStatus(userID string, status string) ([]Payment, error) {
	rows, err := r.DB.Query(
		`SELECT id, merchant, amount, status, created_at, user_id FROM payments WHERE user_id = ? AND status = ? ORDER BY created_at DESC`,
		userID, status,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []Payment
	for rows.Next() {
		var p Payment
		if err := rows.Scan(&p.ID, &p.Merchant, &p.Amount, &p.Status, &p.CreatedAt, &p.UserID); err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, rows.Err()
}