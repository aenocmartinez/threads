package usecase

import (
	"fmt"
	"strings"
	"threads/src/domain"
	"threads/src/infraestructure/middleware"
	"threads/src/shared"
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

func (uc *RegistrarUsuarioUseCase) Execute(name, email, password string) dto.ResponseThreads {
	user, err := uc.userRepo.FindByEmail(email)
	if err != nil {
		return *dto.NewResponseThreads(500, err.Error(), nil)
	}

	if user.Exists() {
		return *dto.NewResponseThreads(409, "Existe un usuario registrado con este email", nil)
	}

	username, err := shared.GenerarUsernameDisponibleDesdeRepositorio(name, uc.userRepo)
	if err != nil {
		return *dto.NewResponseThreads(500, "No se pudo generar un username disponible", nil)
	}

	hashedPassword, err := middleware.HashPassword(password)
	if err != nil {
		return *dto.NewResponseThreads(500, "Ha ocurrido un error con la contraseña", nil)
	}

	user.SetEmail(email)
	user.SetUsername(username)
	user.SetName(name)
	user.SetPassword(hashedPassword)

	if strings.TrimSpace(user.GetPhone()) == "" {
		user.SetPhone("")
	}

	if err := user.Save(); err != nil {
		return *dto.NewResponseThreads(500, fmt.Sprintf("Error al guardar: %v", err), nil)
	}

	return *dto.NewResponseThreads(201, "Usuario registrado exitosamente.", user.ToDTO())
}
