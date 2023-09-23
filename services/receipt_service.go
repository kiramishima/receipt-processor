package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/kiramishima/receipt-processor/domain"
	appErrors "github.com/kiramishima/receipt-processor/pkg/errors"
	ports "github.com/kiramishima/receipt-processor/ports/repository"
	"go.uber.org/zap"
	"strconv"
	"time"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type ReceiptService struct {
	logger     *zap.SugaredLogger
	repository ports.IReceiptRepository
}

func NewReceiptService(repository ports.IReceiptRepository, logger *zap.SugaredLogger) *ReceiptService {
	return &ReceiptService{
		logger:     logger,
		repository: repository,
	}
}

// StoreReceipt Save the points and return the id for consulting
func (svc *ReceiptService) StoreReceipt(ctx context.Context, base *domain.ReceiptBase) (string, error) {
	// Validate
	err := validate.StructCtx(ctx, *base)
	if err != nil {
		svc.logger.Error(err)
		select {
		case <-ctx.Done():
			return "", appErrors.ErrTimeout

		default:

			for _, err := range err.(validator.ValidationErrors) {
				return "", errors.New(fmt.Sprintf("Field: %s, Error: %s\n", err.Field(), err.Tag()))
			}
		}
	}

	// Parse to Receipt
	timeString := fmt.Sprintf("%s %s", base.PurchaseDate, base.PurchaseTime)
	purchaseDt, err := time.Parse("2006-01-02 15:04", timeString)
	if err != nil {
		return "", errors.New(fmt.Sprintf("error parsing timestring: %s", timeString))
	}

	total, err := strconv.ParseFloat(base.Total, 32)
	if err != nil {
		return "", errors.New(fmt.Sprintf("error parsing currency string: %s", base.Total))
	}
	// Items
	var items []*domain.ReceiptItem
	for _, item := range base.Items {
		// Parsing price
		price, err := strconv.ParseFloat(item.Price, 32)
		if err != nil {
			return "", errors.New(fmt.Sprintf("error parsing currency string: %s", item.Price))
		}
		items = append(items, &domain.ReceiptItem{ShortDescription: item.ShortDescription, Price: float32(price)})
	}

	receipt := &domain.Receipt{
		Retailer:   base.Retailer,
		PurchaseDT: purchaseDt,
		Total:      float32(total),
		Items:      items,
	}

	var points = receipt.GetTotalPoints()

	id, err := svc.repository.SaveReceiptPoints(int16(points))
	if err != nil {
		return "", err
	}
	return id, nil
}

// RetrieveReceipt recover points by id
func (svc *ReceiptService) RetrieveReceipt(id string) (*domain.Result, error) {
	svc.logger.Info("ID -> ", id)

	item, err := svc.repository.FindReceiptById(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}
