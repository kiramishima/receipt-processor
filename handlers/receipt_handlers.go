package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/kiramishima/receipt-processor/domain"
	appErrors "github.com/kiramishima/receipt-processor/pkg/errors"
	"github.com/kiramishima/receipt-processor/pkg/utils"
	"github.com/unrolled/render"
	"net/http"

	ports "github.com/kiramishima/receipt-processor/ports/services"
	"go.uber.org/zap"
)

// NewReceiptHandlers creates a instance of auth handlers
func NewReceiptHandlers(r *chi.Mux, logger *zap.SugaredLogger, s ports.IReceiptService, render *render.Render) {
	handler := &ReceiptHandlers{
		logger:   logger,
		service:  s,
		response: render,
	}

	r.Route("/receipts", func(r chi.Router) {
		r.Post("/process", handler.ReceiptProcessHandler)
		r.Get("/{id}/points", handler.ReceiptGetPointsHandler)
	})
}

type ReceiptHandlers struct {
	logger   *zap.SugaredLogger
	service  ports.IReceiptService
	response *render.Render
}

func (h *ReceiptHandlers) ReceiptProcessHandler(w http.ResponseWriter, req *http.Request) {
	var jsonReq = &domain.ReceiptBase{}

	err := utils.ReadJSON(w, req, &jsonReq)

	if err != nil {
		h.logger.Error(err.Error())
		_ = h.response.JSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	h.logger.Info(jsonReq)
	ctx := req.Context()
	id, err := h.service.StoreReceipt(ctx, jsonReq)

	if err != nil {
		h.logger.Error(err.Error())

		select {
		case <-ctx.Done():
			_ = h.response.JSON(w, http.StatusGatewayTimeout, appErrors.ErrTimeout)
		default:
			_ = h.response.JSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return
	}

	if err := h.response.JSON(w, http.StatusOK, map[string]string{"id": id}); err != nil {
		h.logger.Error(err)
		_ = h.response.JSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
}

func (h *ReceiptHandlers) ReceiptGetPointsHandler(w http.ResponseWriter, req *http.Request) {
	receiptID := chi.URLParam(req, "id")

	item, err := h.service.RetrieveReceipt(receiptID)

	if err != nil {
		h.logger.Error(err.Error())
		_ = h.response.JSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	if err := h.response.JSON(w, http.StatusOK, item); err != nil {
		h.logger.Error(err)
		_ = h.response.JSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
}
