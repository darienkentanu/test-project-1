package model

import "time"

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Token      string
	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}

type User_Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
