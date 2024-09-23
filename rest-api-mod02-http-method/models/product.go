package models

// Product model
type Product struct {
	ID          int    `json:"id"`
	ProductName string `json:"productName"`
	Price       int    `json:"price"`
	ItemID      int    `json:"itemID"`
	ItemName    string `json:"itemName"`
}
