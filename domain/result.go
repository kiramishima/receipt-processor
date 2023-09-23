package domain

type Result struct {
	ID     string `json:"-"`
	Points int16  `json:"points"`
}
