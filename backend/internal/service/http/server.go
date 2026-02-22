package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/durianpay/fullstack-boilerplate/internal/api"
	"github.com/durianpay/fullstack-boilerplate/internal/middleware"
	openapi "github.com/durianpay/fullstack-boilerplate/internal/openapigen"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Server struct {
	router http.Handler
}

const (
	readTimeout  = 10
	writeTimeout = 10
	idleTimeout  = 60
)

func NewServer(apiHandler openapi.ServerInterface, openapiYamlPath string) *Server {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://127.0.0.1:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Requested-With"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	r.Use(middleware.Logging)

	r.Options("/*", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// PUBLIC ROUTES
	r.Post("/dashboard/v1/auth/login", apiHandler.PostDashboardV1AuthLogin)

	// PROTECTED ROUTES
	r.Group(func(priv chi.Router) {
		priv.Use(JWTMiddleware)

		// Summary endpoint
		priv.Get("/dashboard/v1/payments/summary", func(w http.ResponseWriter, r *http.Request) {
			if h, ok := apiHandler.(*api.APIHandler); ok {
				h.GetDashboardV1PaymentsSummary(w, r)
				return
			}
			http.Error(w, "internal handler error", http.StatusInternalServerError)
		})

		// Payments list endpoint
		priv.Get("/dashboard/v1/payments", func(w http.ResponseWriter, r *http.Request) {
			query := r.URL.Query()

			var sortParam *openapi.Sort
			if sortVals := query["sort"]; len(sortVals) > 0 && sortVals[0] != "" {
				s := openapi.Sort(sortVals[0])
				sortParam = &s
			}

			var statusParam *string
			if statusVals := query["status"]; len(statusVals) > 0 && statusVals[0] != "" {
				statusParam = &statusVals[0]
			}

			var idParam *string
			if idVals := query["id"]; len(idVals) > 0 && idVals[0] != "" {
				idParam = &idVals[0]
			}

			params := openapi.GetDashboardV1PaymentsParams{
				Sort:   sortParam,
				Status: statusParam,
				Id:     idParam,
			}

			if h, ok := apiHandler.(*api.APIHandler); ok {
				h.GetDashboardV1Payments(w, r, params)
				return
			}
			apiHandler.GetDashboardV1Payments(w, r, params)
		})

		// Review endpoint
		priv.Put("/dashboard/v1/payment/{id}/review", func(w http.ResponseWriter, r *http.Request) {
			id := chi.URLParam(r, "id")
			h, ok := apiHandler.(*api.APIHandler)
			if !ok {
				http.Error(w, "internal handler error", http.StatusInternalServerError)
				return
			}
			if err := h.Payment.Review(id); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"status":"reviewed"}`))
		})
	})

	return &Server{router: r}
}

func (s *Server) Start(addr string) {
	srv := &http.Server{
		Addr:         addr,
		Handler:      s.router,
		ReadTimeout:  readTimeout * time.Second,
		WriteTimeout: writeTimeout * time.Second,
		IdleTimeout:  idleTimeout * time.Second,
	}

	go func() {
		log.Printf("Server listening on %s", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown error: %v", err)
	}
	log.Println("Server stopped ✔")
}

func (s *Server) Routes() http.Handler {
	return s.router
}