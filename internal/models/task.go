package models


type Task struct {
	ID           int        `json:"task_id" gorm:"primaryKey"`
	Title        string     `json:"title"`
	Info         string     `json:"info"`
	Date         string   `json:"date" gorm:"type:date"`
	Status       string     `json:"status"`
}
