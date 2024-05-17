package repositories

import (
	"github.com/aurindo10/internal/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (c *UserRepository) RegisterUser(p *entities.User) (*UserResponseRepository, error) {
	user := &UserDB{
		Id:        p.Id,
		Username:  p.Username,
		Password:  p.Password,
		Name:      p.Name,
		Last_name: p.Last_name,
		Email:     p.Last_name,
	}
	res := c.db.Create(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return &UserResponseRepository{
		Sucess: true,
		Id:     user.Id,
	}, nil
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
