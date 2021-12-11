package model

import "time"

type Transaction struct {
	Id         int     `json:"id"`
	Number     int     `json:"number"`
	PriceTotal float32 `json:"price"`
	CostTotal  float32 `json:"cost"`
	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}

type Transaction_Detail struct {
	Id            int     `json:"id"`
	TransactionID int     `json:"transaction_id"`
	ItemID        int     `json:"item_id"`
	ItemQuantity  int     `json:"item_quantity"`
	ItemPrice     float32 `json:"price"`
	ItemCost      float32 `json:"cost"`
	Created_at    time.Time
	Updated_at    time.Time
	Deleted_at    time.Time
}

type Transaction_Input struct {
	Items []Transaction_Item_Input `json:"item_input"`
}

type Transaction_Item_Input struct {
	Id       int `json:"id"`
	Quantity int `json:"quantity"`
}
