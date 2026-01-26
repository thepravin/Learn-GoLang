package models

type Notes struct {
	Id     int    `json:"id" gorm:"primarykey"`
	Title  string `json:"title" binding:"required"`
	Status bool   `json:"status"`
}
