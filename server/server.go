package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kiramishima/receipt-processor/domain"
	"go.uber.org/zap"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

const (
	certFile                = "ssl/Server.crt"
	keyFile                 = "ssl/Server.pem"
	maxHeaderBytes          = 1 << 20
	ctxTimeout              = 10
	ServerReadHeaderTimeout = 3 * time.Second
)

type Server struct {
	router chi.Router
	logger *zap.SugaredLogger
	cfg    *domain.Configuration
}

func NewServer(cfg *domain.Configuration, logger *zap.SugaredLogger, r *chi.Mux) *Server {
	return &Server{
		router: r,
		logger: logger,
		cfg:    cfg,
	}
}

func (s *Server) Run() error {
	// Create a new http.Server with the specified read header timeout and handler
	var addr = fmt.Sprintf("%s:%d", s.cfg.ServerAddress, s.cfg.Port)

	server := &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: ServerReadHeaderTimeout,
		MaxHeaderBytes:    maxHeaderBytes,
		Handler:           s.router,
	}

	go func() {
		s.logger.Infof("Server is listening on PORT: %d", s.cfg.Port)

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Fatalf("Error starting Server: ", err)
		}
	}()
	return nil
	// graceful shutdown
	// return waitForShutdown(s.logger, server)
}

// waitForShutdown graceful shutdown
func waitForShutdown(logger *zap.SugaredLogger, server *http.Server) error {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig

	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Info("Server Exited Properly")
		logger.Fatal("Failed gracefully")
		logger.Fatal("failed to gracefully shut down server", err)
		return err
	}
	return nil
}

// Module Server Module
var Module = fx.Module("server",
	fx.Provide(func(cfg *domain.Configuration, logger *zap.SugaredLogger, r *chi.Mux) *Server {
		return NewServer(cfg, logger, r)
	}),
)
