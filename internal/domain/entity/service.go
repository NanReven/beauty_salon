package entity

import "time"

type Category struct {
	Id          int
	Title       string
	Description string
	Slug        string
}

type Service struct {
	Id         int
	CategoryId int
	Title      string
	Duration   time.Duration
	Price      float64
}
