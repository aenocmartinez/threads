package usecase

import (
	"threads/src/domain"
	"threads/src/view/dto"
)

type SeguirUsuarioUseCase struct {
	usuarioRepo domain.UserRepository
}

func NewSeguirUsuarioUseCase(usuarioRepo domain.UserRepository) *SeguirUsuarioUseCase {
	return &SeguirUsuarioUseCase{
		usuarioRepo: usuarioRepo,
	}
}

func (uc *SeguirUsuarioUseCase) Execute(miUsuarioID int64, usuarioASeguirID int64) *dto.ResponseThreads {

	miUsuario, _ := uc.usuarioRepo.FindByID(miUsuarioID)
	if !miUsuario.Exists() {
		return dto.NewResponseThreads(404, "Usuario no encontrado", nil)
	}

	usuarioASeguir, _ := uc.usuarioRepo.FindByID(usuarioASeguirID)
	if !usuarioASeguir.Exists() {
		return dto.NewResponseThreads(404, "Usuario a seguir no encontrado", nil)
	}

	exito := miUsuario.SeguirUsuario(usuarioASeguir)
	if !exito {
		return dto.NewResponseThreads(500, "Ha ocurrido un error en el sistema", nil)
	}

	return dto.NewResponseThreads(200, "¡Ahora sigues a este usuario!", nil)
}
