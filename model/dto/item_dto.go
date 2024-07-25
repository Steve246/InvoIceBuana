package dto

type CreateItem struct {
	ItemName  string  `json:"itemName"`
	ItemType  string  `json:"itemType"`
	ItemPrice float64 `json:"itemPrice"`
}
