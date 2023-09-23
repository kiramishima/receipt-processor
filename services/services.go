package services

import (
	repository "github.com/kiramishima/receipt-processor/adapter/db/in-memory"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("services",
	fx.Provide(func(logger *zap.SugaredLogger, receiptRepository *repository.ReceiptRepository) *ReceiptService {
		return NewReceiptService(receiptRepository, logger)
	}),
)
