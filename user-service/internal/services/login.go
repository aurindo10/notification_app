package services

import (
	"github.com/aurindo10/internal/repositories"
	usecases "github.com/aurindo10/internal/usecases/user"
	"gorm.io/gorm"
)

func Login(db *gorm.DB, c *repositories.LoginParams) (*repositories.ResponseParamsLogin, error) {
	repo := repositories.NewUserRepository(db)
	useCase := usecases.NewLoginUseCase(repo)
	res, err := useCase.Execute(c)
	if err != nil {
		return nil, err
	}
	return res, nil
}
