package models



// Task структура представляет собой задачу в вашем списке дел.


type Task struct {
	ID     uint      `json:"id" gorm:"primaryKey"`
	Title  string    `json:"title"`
	Info   string    `json:"info"`
	Date   string    `json:"date" gorm:"type:date"`
	Status string    `json:"status"`
}	

