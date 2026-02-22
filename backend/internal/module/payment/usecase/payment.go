package usecase

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/durianpay/fullstack-boilerplate/internal/module/payment/repository"
)

type Payment struct {
	ID        string
	Merchant  string
	Amount    string
	Status    string
	CreatedAt time.Time
	UserID    string
}

type ListFilter struct {
	Status *string
	Sort   *string
	ID     *string
}

type Summary struct {
	Total      int `json:"total"`
	Completed  int `json:"completed"`
	Processing int `json:"processing"`
	Failed     int `json:"failed"`
}

type Usecase struct {
	repo *repository.Repo
}

func New(repo *repository.Repo) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) ListByUser(userID string) ([]Payment, error) {
	return u.ListByUserFiltered(userID, ListFilter{})
}

func (u *Usecase) ListByUserFiltered(userID string, filter ListFilter) ([]Payment, error) {
	var dbPayments []repository.Payment
	var err error

	if filter.Status != nil && *filter.Status != "" {
		dbPayments, err = u.repo.ListByUserAndStatus(userID, *filter.Status)
	} else {
		dbPayments, err = u.repo.ListByUser(userID)
	}
	if err != nil {
		return nil, err
	}

	var out []Payment
	for _, p := range dbPayments {
		out = append(out, Payment{
			ID:        p.ID,
			Merchant:  p.Merchant,
			Amount:    fmt.Sprintf("%d", p.Amount),
			Status:    p.Status,
			CreatedAt: p.CreatedAt,
			UserID:    p.UserID,
		})
	}

	// ✅ SORTING dengan Tagged Switch
	if filter.Sort != nil && *filter.Sort != "" {
		sortBy := *filter.Sort
		desc := strings.HasPrefix(sortBy, "-")
		key := strings.TrimPrefix(sortBy, "-")

		sort.Slice(out, func(i, j int) bool {
			var less bool

			// Tagged switch on key
			switch key {
			case "amount":
				less = out[i].Amount < out[j].Amount
			case "created_at", "date":
				less = out[i].CreatedAt.Before(out[j].CreatedAt)
			case "merchant":
				less = out[i].Merchant < out[j].Merchant
			case "status":
				less = out[i].Status < out[j].Status
			default:
				return false // unknown field, no sorting
			}

			if desc {
				return !less
			}
			return less
		})
	}

	return out, nil
}

func (u *Usecase) GetSummary(userID string) (*Summary, error) {
	payments, err := u.repo.ListByUser(userID)
	if err != nil {
		return nil, err
	}

	summary := &Summary{Total: len(payments)}

	// Tagged switch untuk counting status
	for _, p := range payments {
		switch p.Status {
		case "completed":
			summary.Completed++
		case "processing":
			summary.Processing++
		case "failed":
			summary.Failed++
		// default: ignored (unknown status)
		}
	}

	return summary, nil
}

func (u *Usecase) Review(paymentID string) error {
	payment, err := u.repo.GetByID(paymentID)
	if err != nil {
		return err
	}

	if payment.Status != "processing" {
		return fmt.Errorf("can only review payments with status 'processing'")
	}

	if err := u.repo.UpdateStatus(paymentID, "completed"); err != nil {
		return err
	}

	return nil
}