package usecase

import (
	"fmt"
	"github.com/durianpay/fullstack-boilerplate/internal/module/payment/repository"
)

type Payment struct {
	ID     string
	Amount string
	Status string
	UserID string
}

type Usecase struct {
	repo *repository.Repo
}

func New(repo *repository.Repo) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) ListByUser(userID string) ([]Payment, error) {
	dbPayments, err := u.repo.ListByUser(userID)
	if err != nil {
		return nil, err
	}

	var out []Payment
	for _, p := range dbPayments {
		out = append(out, Payment{
			ID:     p.ID,
			Amount: fmt.Sprintf("%d", p.Amount),
			Status: p.Status,
			UserID: p.UserID,
		})
	}

	return out, nil
}

// ✅ FIXED: Update status di database
func (u *Usecase) Review(paymentID string) error {
	// Get payment from DB
	payment, err := u.repo.GetByID(paymentID)
	if err != nil {
		return err
	}

	// Only allow reviewing "processing" payments
	if payment.Status != "processing" {
		return fmt.Errorf("can only review payments with status 'processing'")
	}

	// Update status to "completed"
	if err := u.repo.UpdateStatus(paymentID, "completed"); err != nil {
		return err
	}

	return nil
}