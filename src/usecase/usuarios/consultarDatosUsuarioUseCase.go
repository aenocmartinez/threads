package usecase

import (
	"threads/src/domain"
	"threads/src/view/dto"
)

type ConsultarUsuarioUseCase struct {
	userRepo domain.UserRepository
}

func NewConsultarDatosUsuarioUseCase(userRepo domain.UserRepository) *ConsultarUsuarioUseCase {
	return &ConsultarUsuarioUseCase{userRepo: userRepo}
}

func (uc *ConsultarUsuarioUseCase) Execute(usuarioID int64) *dto.ResponseThreads {
	user, _ := uc.userRepo.FindByID(usuarioID)
	if user == nil || !user.Exists() {
		return dto.NewResponseThreads(404, "Usuario no encontrado", nil)
	}

	return dto.NewResponseThreads(200, "Usuario encontrado", user.ToDTO())
}
