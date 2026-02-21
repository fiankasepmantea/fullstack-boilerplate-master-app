package api

import (
	"encoding/json"
	"net/http"

	ah "github.com/durianpay/fullstack-boilerplate/internal/module/auth/handler"
	openapi "github.com/durianpay/fullstack-boilerplate/internal/openapigen"
	"github.com/durianpay/fullstack-boilerplate/internal/middleware"
	payuc "github.com/durianpay/fullstack-boilerplate/internal/module/payment/usecase"
)

type APIHandler struct {
	Auth *ah.AuthHandler
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

	list, _ := h.Payment.ListByUser(userID)

	var out []openapi.Payment
	for _, p := range list {
		out = append(out, openapi.Payment{
			Id:     ptr(p.ID),
			Amount: ptr(p.Amount),
			Status: ptr(p.Status),
		})
	}

	json.NewEncoder(w).Encode(openapi.PaymentListResponse{
		Payments: &out,
	})
}

func ptr(s string) *string { return &s }