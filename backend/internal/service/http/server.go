package http

import (
	"context"
	"log"
	nethttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/durianpay/fullstack-boilerplate/internal/api"
	"github.com/durianpay/fullstack-boilerplate/internal/middleware"
	"github.com/durianpay/fullstack-boilerplate/internal/openapigen"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Server struct {
	router nethttp.Handler
}

const (
	readTimeout  = 10
	writeTimeout = 10
	idleTimeout  = 60
)

func NewServer(apiHandler openapigen.ServerInterface, openapiYamlPath string) *Server {
	r := chi.NewRouter()

	// ✅ GLOBAL middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}))

	r.Use(middleware.Logging)

	// preflight
	r.Options("/*", func(w nethttp.ResponseWriter, r *nethttp.Request) {
		w.WriteHeader(nethttp.StatusOK)
	})

	r.Post("/dashboard/v1/auth/login", apiHandler.PostDashboardV1AuthLogin)

	r.Group(func(priv chi.Router) {
		priv.Use(JWTMiddleware)

		// register OpenAPI routes
		openapigen.HandlerFromMux(apiHandler, priv)

		// review endpoint
		priv.Put("/dashboard/v1/payment/{id}/review", func(w nethttp.ResponseWriter, r *nethttp.Request) {
			id := chi.URLParam(r, "id")

			h := apiHandler.(*api.APIHandler)

			err := h.Payment.Review(id)
			if err != nil {
				nethttp.Error(w, err.Error(), 500)
				return
			}

			w.WriteHeader(nethttp.StatusOK)
			w.Write([]byte(`{"status":"reviewed"}`))
		})
	})

	return &Server{router: r}
}

func (s *Server) Start(addr string) {
	srv := &nethttp.Server{
		Addr:         addr,
		Handler:      s.router,
		ReadTimeout:  readTimeout * time.Second,
		WriteTimeout: writeTimeout * time.Second,
		IdleTimeout:  idleTimeout * time.Second,
	}

	go func() {
		log.Printf("Server listening on %s", addr)
		if err := srv.ListenAndServe(); err != nil && err != nethttp.ErrServerClosed {
			log.Fatalf("listen error: %v", err)
		}
	}()

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("shutdown error: %v", err)
	}

	log.Println("Server stopped ✔")
}

func (s *Server) Routes() nethttp.Handler {
	return s.router
}