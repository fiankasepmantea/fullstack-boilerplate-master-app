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

	filter := payuc.ListFilter{
		Status: params.Status,
		Sort:   params.Sort,
		ID:     params.Id,
	}

	list, err := h.Payment.ListByUserFiltered(userID, filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var out []openapi.Payment
	for _, p := range list {
		id := p.ID
		merchant := p.Merchant
		amt := p.Amount
		st := p.Status
		createdAt := p.CreatedAt

		out = append(out, openapi.Payment{
			Id:        &id,
			Merchant:  &merchant,
			Amount:    &amt,
			Status:    &st,
			CreatedAt: &createdAt,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(openapi.PaymentListResponse{
		Payments: &out,
	})
}

// NEW: Get payment summary
func (h *APIHandler) GetDashboardV1PaymentsSummary(
	w http.ResponseWriter,
	r *http.Request,
) {
	userID, _ := r.Context().Value(middleware.UserIDKey).(string)

	summary, err := h.Payment.GetSummary(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}