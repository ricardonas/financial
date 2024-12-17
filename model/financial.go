package model

import "time"

type Financial struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Value   int64     `json:"value"`
	DueDate time.Time `json:"dueDate"`
	PaidAt  time.Time `json:"paidAt"`
}
