package model

type Todo struct {
	Id     int    `json:"id"`
	Task   string `json:"task"`
	Status bool   `json:"status"`
}
