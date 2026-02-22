package api

import (
	"encoding/json"
	"net/http"

	"github.com/durianpay/fullstack-boilerplate/internal/middleware"
	ah "github.com/durianpay/fullstack-boilerplate/internal/module/auth/handler"
	payuc "github.com/durianpay/fullstack-boilerplate/internal/module/payment/usecase"
	openapi "github.com/durianpay/fullstack-boilerplate/internal/openapigen"
)

type APIHandler struct {
	Auth    *ah.AuthHandler
	Payment *payuc.Usecase
}

var _ openapi.ServerInterface = (*APIHandler)(nil)

func (h *APIHandler) PostDashboardV1AuthLogin(w http.ResponseWriter, r *http.Request) {
	h.Auth.PostDashboardV1AuthLogin(w, r)
}

func (h *APIHandler) GetDashboardV1Payments(
	w http.ResponseWriter,
	r *http.Request,
	params openapi.GetDashboardV1PaymentsParams,
) {
	userID, _ := r.Context().Value(middleware.UserIDKey).(string)

	list, err := h.Payment.ListByUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var out []openapi.Payment
	for _, p := range list {
		id := p.ID
		amt := p.Amount
		st := p.Status

		out = append(out, openapi.Payment{
			Id:     &id,
			Amount: &amt,
			Status: &st,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(openapi.PaymentListResponse{
		Payments: &out,
	})
}