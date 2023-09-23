package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/kiramishima/receipt-processor/domain"
	"github.com/kiramishima/receipt-processor/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
	"testing"
)

func TestReceiptService_StoreReceipt(t *testing.T) {
	logger, _ := zap.NewProduction()
	slogger := logger.Sugar()
	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()
	repo := mocks.NewMockIReceiptRepository(mockCtrl)

	var data = []*domain.ReceiptBase{
		{
			Retailer:     "Target",
			PurchaseDate: "2022-01-01",
			PurchaseTime: "13:01",
			Items: []*domain.ReceiptItemBase{
				{
					ShortDescription: "Mountain Dew 12PK",
					Price:            "6.49",
				}, {
					ShortDescription: "Emils Cheese Pizza",
					Price:            "12.25",
				}, {
					ShortDescription: "Knorr Creamy Chicken",
					Price:            "1.26",
				}, {
					ShortDescription: "Doritos Nacho Cheese",
					Price:            "3.35",
				}, {
					ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
					Price:            "12.00",
				},
			},
			Total: "35.35",
		},
		{
			Retailer:     "",
			PurchaseDate: "2022-01-01",
			PurchaseTime: "13:01",
			Items: []*domain.ReceiptItemBase{
				{
					ShortDescription: "Mountain Dew 12PK",
					Price:            "6.49",
				}, {
					ShortDescription: "Emils Cheese Pizza",
					Price:            "12.25",
				}, {
					ShortDescription: "Knorr Creamy Chicken",
					Price:            "1.26",
				}, {
					ShortDescription: "Doritos Nacho Cheese",
					Price:            "3.35",
				}, {
					ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
					Price:            "12.00",
				},
			},
			Total: "35.35",
		},
		{
			PurchaseDate: "2022-01-01",
			PurchaseTime: "13:01",
			Items: []*domain.ReceiptItemBase{
				{
					ShortDescription: "Mountain Dew 12PK",
					Price:            "6.49",
				}, {
					ShortDescription: "Emils Cheese Pizza",
					Price:            "12.25",
				}, {
					ShortDescription: "Knorr Creamy Chicken",
					Price:            "1.26",
				}, {
					ShortDescription: "Doritos Nacho Cheese",
					Price:            "3.35",
				}, {
					ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
					Price:            "12.00",
				},
			},
			Total: "35.35",
		},
		{
			Retailer:     "Target",
			PurchaseDate: "",
			PurchaseTime: "13:01",
			Items: []*domain.ReceiptItemBase{
				{
					ShortDescription: "Mountain Dew 12PK",
					Price:            "6.49",
				}, {
					ShortDescription: "Emils Cheese Pizza",
					Price:            "12.25",
				}, {
					ShortDescription: "Knorr Creamy Chicken",
					Price:            "1.26",
				}, {
					ShortDescription: "Doritos Nacho Cheese",
					Price:            "3.35",
				}, {
					ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
					Price:            "12.00",
				},
			},
			Total: "35.35",
		},
		{
			Retailer:     "Target",
			PurchaseTime: "13:01",
			Items: []*domain.ReceiptItemBase{
				{
					ShortDescription: "Mountain Dew 12PK",
					Price:            "6.49",
				}, {
					ShortDescription: "Emils Cheese Pizza",
					Price:            "12.25",
				}, {
					ShortDescription: "Knorr Creamy Chicken",
					Price:            "1.26",
				}, {
					ShortDescription: "Doritos Nacho Cheese",
					Price:            "3.35",
				}, {
					ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
					Price:            "12.00",
				},
			},
			Total: "35.35",
		},
		{
			Retailer:     "Target",
			PurchaseDate: "22-01-01",
			PurchaseTime: "13:01",
			Items: []*domain.ReceiptItemBase{
				{
					ShortDescription: "Mountain Dew 12PK",
					Price:            "6.49",
				}, {
					ShortDescription: "Emils Cheese Pizza",
					Price:            "12.25",
				}, {
					ShortDescription: "Knorr Creamy Chicken",
					Price:            "1.26",
				}, {
					ShortDescription: "Doritos Nacho Cheese",
					Price:            "3.35",
				}, {
					ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
					Price:            "12.00",
				},
			},
			Total: "35.35",
		},
		{
			Retailer:     "Target",
			PurchaseDate: "2022-01-01",
			PurchaseTime: "",
			Items: []*domain.ReceiptItemBase{
				{
					ShortDescription: "Mountain Dew 12PK",
					Price:            "6.49",
				}, {
					ShortDescription: "Emils Cheese Pizza",
					Price:            "12.25",
				}, {
					ShortDescription: "Knorr Creamy Chicken",
					Price:            "1.26",
				}, {
					ShortDescription: "Doritos Nacho Cheese",
					Price:            "3.35",
				}, {
					ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
					Price:            "12.00",
				},
			},
			Total: "35.35",
		},
		{
			Retailer:     "Target",
			PurchaseDate: "2022-01-01",
			Items: []*domain.ReceiptItemBase{
				{
					ShortDescription: "Mountain Dew 12PK",
					Price:            "6.49",
				}, {
					ShortDescription: "Emils Cheese Pizza",
					Price:            "12.25",
				}, {
					ShortDescription: "Knorr Creamy Chicken",
					Price:            "1.26",
				}, {
					ShortDescription: "Doritos Nacho Cheese",
					Price:            "3.35",
				}, {
					ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
					Price:            "12.00",
				},
			},
			Total: "35.35",
		},
		{
			Retailer:     "Target",
			PurchaseDate: "2022-01-01",
			PurchaseTime: "25:01",
			Items: []*domain.ReceiptItemBase{
				{
					ShortDescription: "Mountain Dew 12PK",
					Price:            "6.49",
				}, {
					ShortDescription: "Emils Cheese Pizza",
					Price:            "12.25",
				}, {
					ShortDescription: "Knorr Creamy Chicken",
					Price:            "1.26",
				}, {
					ShortDescription: "Doritos Nacho Cheese",
					Price:            "3.35",
				}, {
					ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
					Price:            "12.00",
				},
			},
			Total: "35.35",
		},
		{
			Retailer:     "Target",
			PurchaseDate: "2022-01-01",
			PurchaseTime: "13:01",
			Items:        make([]*domain.ReceiptItemBase, 0),
			Total:        "35.35",
		},
		{
			Retailer:     "Target",
			PurchaseDate: "2022-01-01",
			PurchaseTime: "13:01",
			Total:        "35.35",
		},
		{
			Retailer:     "Target",
			PurchaseDate: "2022-01-01",
			PurchaseTime: "13:01",
			Items: []*domain.ReceiptItemBase{
				{
					ShortDescription: "Mountain Dew 12PK",
					Price:            "6.49",
				}, {
					ShortDescription: "Emils Cheese Pizza",
					Price:            "12.25",
				}, {
					ShortDescription: "Knorr Creamy Chicken",
					Price:            "1.26",
				}, {
					ShortDescription: "Doritos Nacho Cheese",
					Price:            "3.35",
				}, {
					ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
					Price:            "12.00",
				},
			},
			Total: "",
		},
		{
			Retailer:     "Target",
			PurchaseDate: "2022-01-01",
			PurchaseTime: "13:01",
			Items: []*domain.ReceiptItemBase{
				{
					ShortDescription: "Mountain Dew 12PK",
					Price:            "6.49",
				}, {
					ShortDescription: "Emils Cheese Pizza",
					Price:            "12.25",
				}, {
					ShortDescription: "Knorr Creamy Chicken",
					Price:            "1.26",
				}, {
					ShortDescription: "Doritos Nacho Cheese",
					Price:            "3.35",
				}, {
					ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
					Price:            "12.00",
				},
			},
		},
		{
			Retailer:     "Target",
			PurchaseDate: "2022-01-01",
			PurchaseTime: "13:01",
			Items: []*domain.ReceiptItemBase{
				{
					ShortDescription: "Mountain Dew 12PK",
					Price:            "6.49",
				}, {
					ShortDescription: "Emils Cheese Pizza",
					Price:            "12.25",
				}, {
					ShortDescription: "Knorr Creamy Chicken",
					Price:            "1.26",
				}, {
					ShortDescription: "Doritos Nacho Cheese",
					Price:            "3.35",
				}, {
					ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
					Price:            "12.00",
				},
			},
			Total: "abcrf",
		},
	}

	// uids
	var uids = []string{uuid.New().String(), uuid.New().String()}
	gomock.InOrder(
		repo.EXPECT().SaveReceiptPoints(gomock.Any()).Times(1).Return(uids[0], nil),
		repo.EXPECT().SaveReceiptPoints(gomock.Any()).Return("", nil).AnyTimes(),
	)
	svc := NewReceiptService(repo, slogger)

	t.Run("OK", func(t *testing.T) {
		ctx := context.Background()
		data := data[0]
		b, err := svc.StoreReceipt(ctx, data)
		// t.Log("B", b)
		assert.NoError(t, err)
		assert.Equal(t, uids[0], b)
	})
	t.Run("Missing retailer value", func(t *testing.T) {
		ctx := context.Background()
		data := data[1]
		_, err := svc.StoreReceipt(ctx, data)
		// t.Log(b)
		t.Log(err)
		assert.Error(t, err)
		assert.EqualError(t, err, "Field: Retailer, Error: required\n")
	})

	t.Run("Missing retailer key", func(t *testing.T) {
		ctx := context.Background()
		data := data[2]
		_, err := svc.StoreReceipt(ctx, data)
		// t.Log(b)
		t.Log(err)
		assert.Error(t, err)
		assert.EqualError(t, err, "Field: Retailer, Error: required\n")
	})

	t.Run("Empty purchaseDate value", func(t *testing.T) {
		ctx := context.Background()
		data := data[3]
		_, err := svc.StoreReceipt(ctx, data)
		// t.Log(b)
		t.Log(err)
		assert.Error(t, err)
		assert.EqualError(t, err, "Field: PurchaseDate, Error: required\n")
	})

	t.Run("Missing purchaseDate value", func(t *testing.T) {
		ctx := context.Background()
		data := data[4]
		_, err := svc.StoreReceipt(ctx, data)
		// t.Log(b)
		t.Log(err)
		assert.Error(t, err)
		assert.EqualError(t, err, "Field: PurchaseDate, Error: required\n")
	})

	t.Run("Bad purchaseDate", func(t *testing.T) {
		ctx := context.Background()
		data := data[5]
		_, err := svc.StoreReceipt(ctx, data)
		// t.Log(b)
		t.Log(err)
		assert.Error(t, err)
		errmsg := fmt.Sprintf("error parsing timestring: %s %s", data.PurchaseDate, data.PurchaseTime)
		assert.EqualError(t, err, errmsg)
	})

	t.Run("Empty purchaseTime", func(t *testing.T) {
		ctx := context.Background()
		data := data[6]
		_, err := svc.StoreReceipt(ctx, data)
		// t.Log(b)
		t.Log(err)
		assert.Error(t, err)
		assert.EqualError(t, err, "Field: PurchaseTime, Error: required\n")
	})

	t.Run("Missing purchaseDate value", func(t *testing.T) {
		ctx := context.Background()
		data := data[7]
		_, err := svc.StoreReceipt(ctx, data)
		// t.Log(b)
		t.Log(err)
		assert.Error(t, err)
		assert.EqualError(t, err, "Field: PurchaseTime, Error: required\n")
	})

	t.Run("Bad purchaseTime", func(t *testing.T) {
		ctx := context.Background()
		data := data[8]
		_, err := svc.StoreReceipt(ctx, data)
		// t.Log(b)
		t.Log(err)
		assert.Error(t, err)
		errmsg := fmt.Sprintf("error parsing timestring: %s %s", data.PurchaseDate, data.PurchaseTime)
		assert.EqualError(t, err, errmsg)
	})

	t.Run("Ok items", func(t *testing.T) {
		ctx := context.Background()
		data := data[0]
		_, err := svc.StoreReceipt(ctx, data)
		// t.Log(b)
		t.Log(err)
		assert.NoError(t, err)
		assert.Equal(t, len(data.Items), 5)
	})

	t.Run("Missing items values", func(t *testing.T) {
		ctx := context.Background()
		data := data[9]
		_, err := svc.StoreReceipt(ctx, data)
		// t.Log(b)
		t.Log(err)
		assert.Error(t, err)
		assert.EqualError(t, err, "Field: Items, Error: min\n")
	})

	t.Run("Missing items key", func(t *testing.T) {
		ctx := context.Background()
		data := data[10]
		_, err := svc.StoreReceipt(ctx, data)
		// t.Log(b)
		t.Log(err)
		assert.Error(t, err)
		assert.EqualError(t, err, "Field: Items, Error: required\n")
	})

	t.Run("Missing total value", func(t *testing.T) {
		ctx := context.Background()
		data := data[11]
		_, err := svc.StoreReceipt(ctx, data)
		// t.Log(b)
		t.Log(err)
		assert.Error(t, err)
		assert.EqualError(t, err, "Field: Total, Error: required\n")
	})

	t.Run("Missing total key", func(t *testing.T) {
		ctx := context.Background()
		data := data[12]
		_, err := svc.StoreReceipt(ctx, data)
		// t.Log(b)
		t.Log(err)
		assert.Error(t, err)
		assert.EqualError(t, err, "Field: Total, Error: required\n")
	})

	t.Run("Bad total value", func(t *testing.T) {
		ctx := context.Background()
		data := data[13]
		_, err := svc.StoreReceipt(ctx, data)
		// t.Log(b)
		t.Log(err)
		assert.Error(t, err)
		assert.EqualError(t, err, fmt.Sprintf("error parsing currency string: %s", data.Total))
	})
}

func TestReceiptService_StoreReceipt2(t *testing.T) {
	logger, _ := zap.NewProduction()
	slogger := logger.Sugar()
	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()
	repo := mocks.NewMockIReceiptRepository(mockCtrl)

	var uids = []string{uuid.New().String(), uuid.New().String()}
	gomock.InOrder(
		repo.EXPECT().FindReceiptById(gomock.Eq(uids[0])).Times(1).Return(&domain.Result{
			ID:     uids[0],
			Points: 10,
		}, nil),
		repo.EXPECT().FindReceiptById(gomock.Eq(uids[1])).Return(nil, errors.New(fmt.Sprintf("element with id: %s don't found", uids[1]))).AnyTimes(),
	)
	svc := NewReceiptService(repo, slogger)

	t.Run("OK", func(t *testing.T) {
		id := uids[0]
		b, err := svc.RetrieveReceipt(id)
		// t.Log("B", b)
		assert.NoError(t, err)
		assert.Equal(t, id, b.ID)
	})

	t.Run("Not Found", func(t *testing.T) {
		id := uids[1]
		_, err := svc.RetrieveReceipt(id)
		// t.Log("B", b)
		assert.Error(t, err)
	})
}
