package services

import (
	"fmt"
	"ginLearning/04_DB/models"

	"gorm.io/gorm"
)

type NotesService struct {
	db *gorm.DB
}

func (n *NotesService) InitService(database *gorm.DB) {
	n.db = database
	n.db.AutoMigrate(&models.Notes{})
}

func (n *NotesService) GetNotesService() ([]models.Notes, error) {
	var notes []models.Notes
	if err := n.db.Find(&notes).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return notes, nil
}

func (n *NotesService) CreateNotesService(title string, status bool) (*models.Notes, error) {
	notes := &models.Notes{
		Title:  title,
		Status: status,
	}
	if err := n.db.Create(notes).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return notes, nil
}
