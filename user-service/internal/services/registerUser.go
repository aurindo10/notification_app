package services

import (
	"github.com/aurindo10/internal/entities"
	"github.com/aurindo10/internal/repositories"
	usecases "github.com/aurindo10/internal/usecases/user"
	"gorm.io/gorm"
)

func RegisterUserService(db *gorm.DB, c *entities.UserResquest) (*repositories.UserResponseRepository, error) {
	repo := repositories.NewUserRepository(db)
	usecase := usecases.NewRegisterUserUserCase(repo)
	res, err := usecase.Execute(c)
	if err != nil {
		return nil, err
	}
	return res, nil
}
