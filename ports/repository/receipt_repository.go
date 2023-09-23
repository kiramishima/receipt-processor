package repository

import "github.com/kiramishima/receipt-processor/domain"

type IReceiptRepository interface {
	SaveReceiptPoints(points int16) (string, error)
	FindReceiptById(id string) (*domain.Result, error)
}
