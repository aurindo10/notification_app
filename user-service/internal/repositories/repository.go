package repositories

import (
	"github.com/aurindo10/internal/entities"
	"github.com/google/uuid"
)

type UserResponseRepository struct {
	Sucess bool
	Id     uuid.UUID
}

type Repository interface {
	RegisterUser(p *entities.User) (*UserResponseRepository, error)
}

type UserDB struct {
	Id        uuid.UUID `gorm:"primaryKey;column:id"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	Name      string    `gorm:"column:name"`
	Last_name *string   `gorm:"column:last_name"`
	Email     *string   `gorm:"column:email"`
}
