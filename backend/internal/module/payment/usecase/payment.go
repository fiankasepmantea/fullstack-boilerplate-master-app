package usecase

import (
	"sort"

	"github.com/durianpay/fullstack-boilerplate/internal/module/payment/repository"
)

type Usecase struct {
	repo *repository.Repo
}

func New(repo *repository.Repo) *Usecase {
	return &Usecase{repo: repo}
}

type ListFilter struct {
	Status *string
	ID     *string
	Sort   *string
}

func (u *Usecase) ListByUserFiltered(
	userID string,
	f ListFilter,
) ([]repository.Payment, error) {

	list, err := u.repo.ListByUser(userID)
	if err != nil {
		return nil, err
	}

	var out []repository.Payment

	for _, p := range list {

		// ✅ STATUS FILTER
		if f.Status != nil && *f.Status != "" && p.Status != *f.Status {
			continue
		}

		// ✅ ID FILTER
		if f.ID != nil && *f.ID != "" && p.ID != *f.ID {
			continue
		}

		out = append(out, p)
	}

	// ✅ SORT
	if f.Sort != nil {
		switch *f.Sort {
		case "amount_asc":
			sort.Slice(out, func(i, j int) bool {
				return out[i].Amount < out[j].Amount
			})
		case "amount_desc":
			sort.Slice(out, func(i, j int) bool {
				return out[i].Amount > out[j].Amount
			})
		}
	}

	return out, nil
}

func (u *Usecase) Review(id string) error {
	p, err := u.repo.FindByID(id)
	if err != nil {
		return err
	}

	p.Status = "success"
	return u.repo.Update(p)
}