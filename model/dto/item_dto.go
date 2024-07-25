package dto

type CreateItem struct {
	ItemName  string  `json:"itemName"`
	ItemType  string  `json:"itemType"`
	ItemPrice float64 `json:"itemPrice"`
}

type DisplayItem struct {
	ItemId             string  `json:"itemId"`
	ItemName           string  `json:"itemName"`
	ItemType           string  `json:"itemType"`
	ItemPrice          float64 `json:"itemPrice"`
	ItemPriceFormatted string  `json:"itemPriceFormatted"`
}
