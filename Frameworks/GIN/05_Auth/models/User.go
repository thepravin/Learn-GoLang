package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" binding:"required" gorm:"not null"`
	Email    string `json:"email" binding:"required" gorm:"uniqueIndex;not null"`
	Password string `json:"password" binding:"required" gorm:"not null"`
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (User) TableName() string {
	return "users"
}

type Notes struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title" gorm:"not null"`
	Status bool   `json:"status" gorm:"default:false"`
}

func (Notes) TableName() string {
	return "notes"
}
