package usecase

import (
	"threads/src/domain"
	"threads/src/view/dto"
)

type DejarDeSeguirUsuarioUseCase struct {
	usuaRepo domain.UserRepository
}

func NewDejarDeSeguirUsuarioUseCase(usuaRepo domain.UserRepository) *DejarDeSeguirUsuarioUseCase {
	return &DejarDeSeguirUsuarioUseCase{
		usuaRepo: usuaRepo,
	}
}

func (uc *DejarDeSeguirUsuarioUseCase) Execute(miUsuarioID int64, usuarioDejarSeguirID int64) *dto.ResponseThreads {

	miUsuario, _ := uc.usuaRepo.FindByID(miUsuarioID)
	if !miUsuario.Exists() {
		return dto.NewResponseThreads(404, "Usuario no encontrado", nil)
	}

	usuarioDejarSeguir, _ := uc.usuaRepo.FindByID(usuarioDejarSeguirID)
	if !usuarioDejarSeguir.Exists() {
		return dto.NewResponseThreads(404, "Usuario a seguir no encontrado", nil)
	}

	exito := miUsuario.DejarDeSeguirUsuario(usuarioDejarSeguir)
	if !exito {
		return dto.NewResponseThreads(500, "Ha ocurrido un error en el sistema", nil)
	}

	return dto.NewResponseThreads(200, "Ya no sigues a este usuario.", nil)
}
