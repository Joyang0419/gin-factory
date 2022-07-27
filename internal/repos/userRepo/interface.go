package userRepo

import (
	"aCupOfGin/internal/entities"
)

type InterfaceUserRepo interface {
	CreateUser(
		vendor string,
		account string,
		accountType string,
		hashedPassword string,
		name string) (bool, error)
	GetAllUsers() []entities.UserEntity
	DeleteUser(id int) (bool, error)
	GetUser(id int) *entities.UserEntity
	UpdateUser(id int, name string) (bool, error)
}
