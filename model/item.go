package model

import "time"

type Item struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Price      float32 `json:"price"`
	Cost       float32 `json:"cost"`
	Created_at time.Time
	Updated_at time.Time
}
