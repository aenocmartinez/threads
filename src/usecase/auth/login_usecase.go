package auth

import (
	"errors"
	"threads/src/infraestructure/di"
	"threads/src/infraestructure/middleware"
	"threads/src/view/dto"
)

type LoginUseCase struct{}

func (uc *LoginUseCase) Execute(username, password string) (*dto.UserDTO, error) {

	userRepo := di.GetContainer().GetUserRepository()

	user, err := userRepo.FindUserLogin(username)
	if err != nil {
		return nil, err
	}

	if !user.Exists() {
		return nil, errors.New("usuario no encontrado")
	}

	if !middleware.VerifyPassword(user.GetPassword(), password) {
		return nil, errors.New("contraseña incorrecta")
	}

	// Generar el token JWT con el secreto del usuario
	token, err := middleware.GenerateToken(user.GetID(), user.GetUsername())
	if err != nil {
		return nil, errors.New("error al generar el token")
	}

	userDTO := user.ToDTO()
	userDTO.SessionToken = token

	return userDTO, nil
}
