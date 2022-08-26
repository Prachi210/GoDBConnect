package model

type Product struct {
	ID              int    `json:"id"`
	Model           string `json:"model"`
	ItemType        string `json:"itemType"`
	Category        string `json:"category"`
	Price           string `json:"price"`
	TotalProducts   int    `json:"totalProducts"`
	PerProductCount int    `json:"perProductCount"`
}

// set ID as 0 for EOF
// send dummy prod as EOF
