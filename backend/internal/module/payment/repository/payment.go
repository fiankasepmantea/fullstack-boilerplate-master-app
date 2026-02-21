package repository

import (
	"database/sql"
)

type Payment struct {
	ID     string
	Amount int64
	Status string
	UserID string
}

type Repo struct {
	DB *sql.DB
}

func New(db *sql.DB) *Repo {
	return &Repo{DB: db}
}

func (r *Repo) ListByUser(userID string) ([]Payment, error) {
	rows, err := r.DB.Query(
		`SELECT id, amount, status, user_id FROM payments WHERE user_id=?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []Payment
	for rows.Next() {
		var p Payment
		rows.Scan(&p.ID, &p.Amount, &p.Status, &p.UserID)
		out = append(out, p)
	}

	return out, nil
}