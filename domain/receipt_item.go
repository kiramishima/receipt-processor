package domain

type ReceiptItem struct {
	ShortDescription string  `json:"shortDescription" validate:"required"`
	Price            float32 `json:"price" validate:"required"`
}
