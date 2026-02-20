package api

import (
	"encoding/json"
	"net/http"

	ah "github.com/durianpay/fullstack-boilerplate/internal/module/auth/handler"
	openapi "github.com/durianpay/fullstack-boilerplate/internal/openapigen"
	shttp "github.com/durianpay/fullstack-boilerplate/internal/service/http"
)

type APIHandler struct {
	Auth *ah.AuthHandler
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
	userID, _ := r.Context().Value(shttp.UserIDKey).(string)
	_ = userID

	payments := []openapi.Payment{
		{Id: ptr("pay_001"), Amount: ptr("100000"), Status: ptr("pending")},
		{Id: ptr("pay_002"), Amount: ptr("250000"), Status: ptr("success")},
	}

	resp := openapi.PaymentListResponse{
		Payments: &payments,
	}

	json.NewEncoder(w).Encode(resp)
}

func ptr(s string) *string { return &s }