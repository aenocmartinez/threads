package usecase

import (
	"threads/src/domain"
	"threads/src/infraestructure/middleware"
	"threads/src/view/dto"
)

type RegistrarUsuarioUseCase struct {
	userRepo domain.UserRepository
}

func NewRegistrarUsuarioUseCase(userRepo domain.UserRepository) *RegistrarUsuarioUseCase {
	return &RegistrarUsuarioUseCase{
		userRepo: userRepo,
	}
}

func (uc *RegistrarUsuarioUseCase) Execute(username, name, email, password string) dto.ResponseThreads {

	user, err := uc.userRepo.FindByEmail(email)
	if err != nil {
		return *dto.NewResponseThreads(500, err.Error(), nil)
	}

	if user.Exists() {
		return *dto.NewResponseThreads(409, "Existe un usuario registrado con este email", nil)
	}

	user, err = uc.userRepo.FindByUsername(username)
	if err != nil {
		return *dto.NewResponseThreads(500, err.Error(), nil)
	}

	if user.Exists() {
		return *dto.NewResponseThreads(409, "Existe un usuario registrado con este username", nil)
	}

	hassPassword, err := middleware.HashPassword(password)
	if err != nil {
		return *dto.NewResponseThreads(500, "Ha ocurrido un error con la contraseña", nil)
	}

	user.SetEmail(email)
	user.SetUsername(username)
	user.SetName(name)
	user.SetPassword(hassPassword)

	err = user.Save()
	if err != nil {
		return *dto.NewResponseThreads(500, "Ha ocurrido un error en el sistema durante la creación del usuario", nil)
	}

	return *dto.NewResponseThreads(201, "Usuario registrado exitosamente.", user.ToDTO())
}
