package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/durianpay/fullstack-boilerplate/internal/middleware"
	"github.com/durianpay/fullstack-boilerplate/internal/openapigen"
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

// NewServer make router chi with CORS & JWT middleware
func NewServer(apiHandler openapigen.ServerInterface, openapiYamlPath string) *Server {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))

	// Logging middleware
	r.Use(middleware.Logging)

	// Public routes (login, register, dll)
	r.Group(func(pub chi.Router) {
		openapigen.HandlerFromMux(apiHandler, pub)
	})

	// Protected routes (JWT)
	r.Group(func(priv chi.Router) {
		priv.Use(JWTMiddleware) // pakai JWT middleware
		priv.Get("/dashboard/v1/payments", func(w http.ResponseWriter, r *http.Request) {
			apiHandler.GetDashboardV1Payments(w, r, openapigen.GetDashboardV1PaymentsParams{})
		})
	})

	return &Server{
		router: r,
	}
}

// Start server with graceful shutdown
func (s *Server) Start(addr string) {
	service := &http.Server{
		Addr:         addr,
		Handler:      s.router,
		ReadTimeout:  readTimeout * time.Second,
		WriteTimeout: writeTimeout * time.Second,
		IdleTimeout:  idleTimeout * time.Second,
	}

	go func() {
		log.Printf("Server listening on %s", addr)
		if err := service.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen error: %v", err)
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down gracefully...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := service.Shutdown(ctx); err != nil {
		log.Fatalf("Forced shutdown: %v", err)
	}

	log.Println("Server stopped cleanly ✔")
}

// Routes return router
func (s *Server) Routes() http.Handler {
	return s.router
}