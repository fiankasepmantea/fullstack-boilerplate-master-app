package usecase

import (
	"sort"
	"strings"

	"github.com/durianpay/fullstack-boilerplate/internal/module/payment/repository"
)

type Payment struct {
	ID     string
	Amount string
	Status string
	UserID string
}

type ListFilter struct {
	Status *string
	Sort   *string
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
	// Mock data - SESUAI OPENAPI SPEC (processing, completed, failed)
	allPayments := []Payment{
		{ID: "pay_001", Amount: "100000", Status: "processing", UserID: userID},
		{ID: "pay_002", Amount: "250000", Status: "completed", UserID: userID},
		{ID: "pay_003", Amount: "50000", Status: "failed", UserID: userID},
		{ID: "pay_004", Amount: "500000", Status: "processing", UserID: userID},
	}

	// ✅ FILTER BY STATUS
	if filter.Status != nil && *filter.Status != "" {
		var filtered []Payment
		for _, p := range allPayments {
			if p.Status == *filter.Status {
				filtered = append(filtered, p)
			}
		}
		allPayments = filtered
	}

	// ✅ SORT BY AMOUNT
	if filter.Sort != nil && *filter.Sort != "" {
		sortBy := *filter.Sort
		desc := strings.HasPrefix(sortBy, "-")
		key := strings.TrimPrefix(sortBy, "-")

		sort.Slice(allPayments, func(i, j int) bool {
			// Simple string compare for amount
			if key == "amount" {
				if desc {
					return allPayments[i].Amount > allPayments[j].Amount
				}
				return allPayments[i].Amount < allPayments[j].Amount
			}
			return false
		})
	}

	return allPayments, nil
}

func (u *Usecase) Review(paymentID string) error {
	// Mock: just validate ID
	if paymentID == "" {
		return nil
	}
	// In real app: update status in DB via repository
	return nil
}