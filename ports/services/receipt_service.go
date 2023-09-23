package services

import (
	"context"
	"github.com/kiramishima/receipt-processor/domain"
)

type IReceiptService interface {
	StoreReceipt(ctx context.Context, base *domain.ReceiptBase) (string, error)
	RetrieveReceipt(id string) (*domain.Result, error)
}
