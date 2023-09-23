package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/kiramishima/receipt-processor/services"
	"github.com/unrolled/render"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("handlers",
	fx.Invoke(func(r *chi.Mux, logger *zap.SugaredLogger, svc *services.ReceiptService, render *render.Render) {
		NewReceiptHandlers(r, logger, svc, render)
	}),
)
