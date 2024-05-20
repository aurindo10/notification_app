package repositories

import (
	"github.com/aurindo10/internal/entities"
	"github.com/google/uuid"
)

type UserResponseRepository struct {
	Sucess bool      `json:"sucess"`
	Id     uuid.UUID `json:"id"`
}

type Repository interface {
	RegisterUser(p *entities.User) (*UserResponseRepository, error)
}

type UserDB struct {
	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	Name      string    `gorm:"column:name"`
	Last_name *string   `gorm:"column:last_name"`
	Email     *string   `gorm:"column:email"`
}
