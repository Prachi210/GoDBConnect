package model

type ConsumerProduct struct {
	ID              int    `json:"id"`
	Model           string `json:"model"`
	ItemType        string `json:"itemType"`
	Category        string `json:"category"`
	Price           string `json:"price"`
	TotalProducts   int    `json:"totalProducts"`
	PerProductCount int    `json:"perProductCount"`
}
