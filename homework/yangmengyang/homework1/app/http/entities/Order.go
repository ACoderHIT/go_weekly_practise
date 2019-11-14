package entities

import "time"

type Order struct {
	Id   int64  `json:"id" example:"1"`
	Name string `json:"name" example:"snow"`
	Date time.Time `json:"date" example:"2019-10-31 11:00:00"`
}