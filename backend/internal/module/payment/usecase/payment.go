package usecase

type Payment struct {
	ID     string
	Amount string
	Status string
	UserID string
}

type Usecase struct{}

func New() *Usecase {
	return &Usecase{}
}

func (u *Usecase) ListByUser(userID string) ([]Payment, error) {
	return []Payment{
		{"pay_001", "100000", "pending", userID},
		{"pay_002", "250000", "success", userID},
	}, nil
}