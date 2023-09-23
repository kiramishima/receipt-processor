package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var testCases = map[struct {
	TotalWords       int
	TotalCents       int
	Total25          int
	Total2Items      int
	TotalItemDesc    int
	TotalOddDay      int
	TotalBetweenTime int
	TotalPoints      int
}]Receipt{
	{TotalWords: 6, TotalCents: 0, Total25: 0, Total2Items: 10, TotalItemDesc: 6, TotalOddDay: 6, TotalBetweenTime: 0, TotalPoints: 28}: {
		Retailer:   "Target",
		PurchaseDT: time.Date(2022, 1, 1, 13, 1, 0, 0, time.UTC),
		Items: []*ReceiptItem{
			{
				ShortDescription: "Mountain Dew 12PK",
				Price:            6.49,
			}, {
				ShortDescription: "Emils Cheese Pizza",
				Price:            12.25,
			}, {
				ShortDescription: "Knorr Creamy Chicken",
				Price:            1.26,
			}, {
				ShortDescription: "Doritos Nacho Cheese",
				Price:            3.35,
			}, {
				ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
				Price:            12.00,
			},
		},
		Total: 35.35,
	},
	{TotalWords: 7, TotalCents: 0, Total25: 25, Total2Items: 0, TotalItemDesc: 0, TotalOddDay: 0, TotalBetweenTime: 0, TotalPoints: 32}: {
		Retailer:   "Targets",
		PurchaseDT: time.Date(2022, 1, 2, 13, 13, 0, 0, time.UTC),
		Items: []*ReceiptItem{
			{
				ShortDescription: "Pepsi - 12-oz",
				Price:            1.25,
			},
		},
		Total: 1.25,
	},
	{TotalWords: 9, TotalCents: 0, Total25: 0, Total2Items: 5, TotalItemDesc: 1, TotalOddDay: 0, TotalBetweenTime: 0, TotalPoints: 15}: {
		Retailer:   "Walgreens",
		PurchaseDT: time.Date(2022, 1, 2, 8, 13, 0, 0, time.UTC),
		Items: []*ReceiptItem{
			{
				ShortDescription: "Pepsi - 12-oz",
				Price:            1.25,
			}, {
				ShortDescription: "Dasani",
				Price:            1.40,
			},
		},
		Total: 2.65,
	},
	{TotalWords: 14, TotalCents: 50, Total25: 25, Total2Items: 10, TotalItemDesc: 0, TotalOddDay: 0, TotalBetweenTime: 10, TotalPoints: 109}: {
		Retailer:   "M&M Corner Market",
		PurchaseDT: time.Date(2022, 3, 20, 14, 33, 0, 0, time.UTC),
		Items: []*ReceiptItem{
			{
				ShortDescription: "Gatorade",
				Price:            2.25,
			}, {
				ShortDescription: "Gatorade",
				Price:            2.25,
			},
			{
				ShortDescription: "Gatorade",
				Price:            2.25,
			},
			{
				ShortDescription: "Gatorade",
				Price:            2.25,
			},
		},
		Total: 9.00,
	},
}

func TestReceipt_GetPointsCountingRetailerName(t *testing.T) {
	for key, item := range testCases {
		var receipt = item
		var total = receipt.GetPointsCountingRetailerName()
		assert.Equal(t, key.TotalWords, total)
	}
}

func TestReceipt_GetPointsIfTotalRoundWithNoCents(t *testing.T) {
	for key, item := range testCases {
		var receipt = item
		var total = receipt.GetPointsIfTotalRoundWithNoCents()

		assert.Equal(t, key.TotalCents, total)
	}
}

func TestReceipt_GetPointsIfTotalIsMultipleOf25Cents(t *testing.T) {
	for key, item := range testCases {
		var receipt = item
		t.Log(item)
		var total = receipt.GetPointsIfTotalIsMultipleOf25Cents()
		t.Log(total)
		assert.Equal(t, key.Total25, total)
	}
}

func TestReceipt_Get5PointForEvery2Items(t *testing.T) {
	for key, item := range testCases {
		var receipt = item
		var total = receipt.Get5PointForEvery2Items()

		assert.Equal(t, key.Total2Items, total)
	}
}

func TestReceipt_GetPointsFromItemsDescription(t *testing.T) {
	for key, item := range testCases {
		var receipt = item
		var total = receipt.GetPointsFromItemsDescription()

		assert.Equal(t, key.TotalItemDesc, total)
	}
}

func TestReceipt_GetPointsDayIsOdd(t *testing.T) {
	for key, item := range testCases {
		var receipt = item
		var total = receipt.GetPointsDayIsOdd()

		assert.Equal(t, key.TotalOddDay, total)
	}
}

func TestReceipt_GetPointsBetweenTime(t *testing.T) {
	for key, item := range testCases {
		var receipt = item
		var total = receipt.GetPointsBetweenTime()

		assert.Equal(t, key.TotalBetweenTime, total)
	}
}

func TestReceipt_GetTotalPoints(t *testing.T) {
	for key, item := range testCases {
		var receipt = item
		var total = receipt.GetTotalPoints()

		assert.Equal(t, key.TotalPoints, total)
	}
}
