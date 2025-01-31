package services

import (
	"data-management-app/db"
	"data-management-app/models"
)

type UserService struct{}

func (s *UserService) CreateUser(user *models.User) error {
	return db.CreateUser(user)
}

func (s *UserService) GetUser(id int) (*models.User, error) {
	return db.GetUser(id)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return db.UpdateUser(user)
}

func (s *UserService) DeleteUser(id int) error {
	return db.DeleteUser(id)
}
