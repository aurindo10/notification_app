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
		Username:  p.Username,
		Password:  p.Password,
		Name:      p.Name,
		Last_name: p.Last_name,
		Email:     p.Email,
	}
	res := c.db.Create(&user)
	if res.Error != nil {
		println(res.Error.Error())
		return nil, res.Error
	}
	return &UserResponseRepository{
		Sucess: true,
		Id:     user.Id,
	}, nil
}
func (c *UserRepository) IsEmailAlreadyExists(email *string) (*UserDB, error) {
	var user UserDB
	if err := c.db.Where("email = ?", email).First(&user).Error; err != nil {
		println(err.Error())
		return nil, err
	}
	return &user, nil
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
