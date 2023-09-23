package domain

type ReceiptBase struct {
	Retailer     string             `json:"retailer,omitempty" validate:"required"`
	PurchaseDate string             `json:"purchaseDate,omitempty" validate:"required"`
	PurchaseTime string             `json:"purchaseTime,omitempty" validate:"required"`
	Total        string             `json:"total,omitempty" validate:"required"`
	Items        []*ReceiptItemBase `json:"items,omitempty" validate:"required,min=1,dive,required"`
}
