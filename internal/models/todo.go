package db

type Todo struct {
	ID uint `json:"id" gorm:"primary_key"`
	Content string `json:"content"`
	IsComplete bool `json:"isComplete"`
}