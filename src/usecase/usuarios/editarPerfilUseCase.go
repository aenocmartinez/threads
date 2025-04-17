package usecase

import (
	"threads/src/domain"
	"threads/src/view/dto"
)

type EditarPerfilUseCase struct {
	userRepo domain.UserRepository
}

func NewEditarPerfilUseCase(userRepo domain.UserRepository) *EditarPerfilUseCase {
	return &EditarPerfilUseCase{userRepo: userRepo}
}

func (uc *EditarPerfilUseCase) Execute(input dto.EditarPerfilDTO) *dto.ResponseThreads {

	user, _ := uc.userRepo.FindByID(input.ID)
	if user == nil || !user.Exists() {
		return dto.NewResponseThreads(404, "Usuario no encontrado", nil)
	}

	if input.Username != user.GetUsername() {
		existeUsername, err := uc.userRepo.ExistsUsername(input.Username)
		if err != nil {
			return dto.NewResponseThreads(500, "Error verificando username", nil)
		}
		if existeUsername {
			return dto.NewResponseThreads(409, "El nombre de usuario ya está en uso", nil)
		}
	}

	user.SetName(input.Name)
	// user.SetUsername(input.Username)
	// user.SetEmail(input.Email)
	user.SetPhone(input.Phone)
	user.SetAvatar(input.Avatar)
	user.SetDescription(input.Description)

	err := user.Update()
	if err != nil {
		return dto.NewResponseThreads(500, "No se pudo actualizar el perfil", nil)
	}

	return dto.NewResponseThreads(200, "Perfil actualizado correctamente", user.ToDTO())
}
