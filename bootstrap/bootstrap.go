package bootstrap

import (
	"context"
	repository "github.com/kiramishima/receipt-processor/adapter/db/in-memory"
	"github.com/kiramishima/receipt-processor/handlers"
	"github.com/kiramishima/receipt-processor/services"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kiramishima/receipt-processor/config"
	"github.com/kiramishima/receipt-processor/server"
	"github.com/unrolled/render"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func bootstrap(
	lifecycle fx.Lifecycle,
	logger *zap.SugaredLogger,
	server *server.Server,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				logger.Info("Starting API")
				_ = server.Run()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				return logger.Sync()
			},
		},
	)
}

var Module = fx.Options(
	config.Module,
	fx.Provide(func() *chi.Mux {
		var r = chi.NewRouter()
		r.Use(middleware.Timeout(60 * time.Second))
		r.Use(middleware.RequestID)
		r.Use(middleware.RealIP)
		r.Use(middleware.Recoverer)
		r.Use(middleware.Logger)
		r.Use(middleware.Compress(5))
		return r
	}),
	fx.Provide(func() *render.Render {
		return render.New()
	}),
	server.Module,
	repository.Module,
	services.Module,
	handlers.Module,
	fx.Invoke(bootstrap),
)
