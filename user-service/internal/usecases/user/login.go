package usecases

import (
	"github.com/aurindo10/internal/repositories"
	"github.com/aurindo10/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type LoginUseCase struct {
	repository repositories.Repository
}

func (r *LoginUseCase) Execute(p *repositories.LoginParams) (*repositories.ResponseParamsLogin, error) {
	res, error := r.repository.GetUser(p)
	if error != nil {
		return nil, error
	}
	password := []byte(*p.Password)

	error = bcrypt.CompareHashAndPassword([]byte(res.Password), password)
	if error != nil {
		return nil, error
	}
	token, error := utils.GenerateToken(*res.Email)
	if error != nil {
		return nil, error
	}
	return &repositories.ResponseParamsLogin{
		Token: token,
	}, nil
}

func NewLoginUseCase(repository repositories.Repository) *LoginUseCase {
	return &LoginUseCase{
		repository: repository,
	}
}
