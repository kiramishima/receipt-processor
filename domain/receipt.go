package domain

import (
	"log"
	"math"
	"regexp"
	"strings"
	"time"
)

type Receipt struct {
	Retailer   string         `json:"retailer,omitempty"`
	PurchaseDT time.Time      `json:"purchaseTime,omitempty"`
	Total      float32        `json:"total,omitempty"`
	Items      []*ReceiptItem `json:"items,omitempty"`
}

// GetPointsCountingRetailerName returns 1 point by alphanumeric characters in the retailer name
func (r *Receipt) GetPointsCountingRetailerName() int {
	var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	var retailName = strings.ReplaceAll(nonAlphanumericRegex.ReplaceAllString(r.Retailer, ""), " ", "")
	log.Println("GetPointsCountingRetailerName -> ", len(retailName)*1)
	return len(retailName) * 1
}

// GetPointsIfTotalRoundWithNoCents returns 50 points if the total is rounded amount without cents
func (r *Receipt) GetPointsIfTotalRoundWithNoCents() int {
	var points = 0
	if math.Ceil(float64(r.Total)) == float64(r.Total) {
		points = 50
	}
	log.Println("GetPointsIfTotalRoundWithNoCents -> ", points)
	return points
}

// GetPointsIfTotalIsMultipleOf25Cents returns 25 points if the amount is multiple of 0.25
func (r *Receipt) GetPointsIfTotalIsMultipleOf25Cents() int {
	var points = 0
	if math.Mod(float64(r.Total), 0.25) == 0 {
		points = 25
	}
	log.Println("GetPointsIfTotalIsMultipleOf25Cents -> ", math.Mod(float64(r.Total), 0.25))
	log.Println("GetPointsIfTotalIsMultipleOf25Cents -> ", points)
	return points
}

// Get5PointForEvery2Items returns 5 points by every 2 items on the receipt
func (r *Receipt) Get5PointForEvery2Items() int {
	var totalItems = len(r.Items)
	var totalPoints = 0
	for i := 2; i <= totalItems; i += 2 {
		totalPoints += 5
	}
	log.Println("Get5PointForEvery2Items -> ", totalPoints)
	return totalPoints
}

// GetPointsFromItemsDescription returns points if item description is a multiple of 3, then it multiply the price by 0.2 and rounded by the nearest integer
func (r *Receipt) GetPointsFromItemsDescription() int {
	var totalPoints = 0
	for _, item := range r.Items {
		var txt = strings.TrimSpace(item.ShortDescription)
		if math.Mod(float64(len(txt)), 3.0) == 0 {
			var op = item.Price * 0.2
			// log.Println(op, " -> ", math.Ceil(float64(op)))
			totalPoints += int(math.Ceil(float64(op)))
		}
	}
	log.Println("GetPointsFromItemsDescription -> ", totalPoints)
	return totalPoints
}

// GetPointsDayIsOdd returns 6 points if the day in the purchase date is odd
func (r *Receipt) GetPointsDayIsOdd() int {
	var totalPoints = 0
	// log.Println("GetPointsDayIsOdd -> ", math.Mod(float64(r.PurchaseDT.Day()), 2.0))

	if math.Mod(float64(r.PurchaseDT.Day()), 2.0) != 0 {
		totalPoints = 6
	}
	log.Println("GetPointsDayIsOdd -> ", totalPoints)
	return totalPoints
}

// GetPointsBetweenTime returns 10 points if the time of purchase is after 2:00pm and before 4:00pm
func (r *Receipt) GetPointsBetweenTime() int {
	var totalPoints = 0
	var now = r.PurchaseDT
	var t1 = time.Date(now.Year(), now.Month(), now.Day(), 14, 0, 0, 0, time.UTC)
	var t2 = time.Date(now.Year(), now.Month(), now.Day(), 16, 0, 0, 0, time.UTC)

	if r.PurchaseDT.After(t1) && r.PurchaseDT.Before(t2) {
		totalPoints = 10
	}
	log.Println("GetPointsBetweenTime -> ", totalPoints)
	return totalPoints
}

// GetTotalPoints returns all collected points
func (r *Receipt) GetTotalPoints() int {
	var points = 0
	points += r.GetPointsCountingRetailerName()
	points += r.GetPointsIfTotalRoundWithNoCents()
	points += r.GetPointsIfTotalIsMultipleOf25Cents()
	points += r.Get5PointForEvery2Items()
	points += r.GetPointsFromItemsDescription()
	points += r.GetPointsDayIsOdd()
	points += r.GetPointsBetweenTime()
	log.Println("GetTotalPoints -> ", points)
	return points
}
