package in_memory

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/kiramishima/receipt-processor/domain"
	"github.com/samber/lo"
)

func NewReceiptRepository() *ReceiptRepository {
	return &ReceiptRepository{
		records: make([]*domain.Result, 0),
	}
}

type ReceiptRepository struct {
	records []*domain.Result
}

func (repo *ReceiptRepository) SaveReceiptPoints(points int16) (string, error) {
	// Generate ID
	var uid = uuid.New()
	// Insert record
	repo.records = append(repo.records, &domain.Result{ID: uid.String(), Points: points})
	return uid.String(), nil
}

func (repo *ReceiptRepository) FindReceiptById(id string) (*domain.Result, error) {
	// Find or return empty
	item := lo.FindOrElse(repo.records, nil, func(i *domain.Result) bool {
		return i.ID == id
	})

	if item == nil {
		return nil, errors.New(fmt.Sprintf("element with id: %s don't found", id))
	}
	return item, nil
}
