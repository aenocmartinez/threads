package usecase

import (
	"threads/src/domain"
	"threads/src/view/dto"
)

type DarMeGustaComentarioUseCase struct {
	comentarioRepo domain.ComentarioRepository
	userRepo       domain.UserRepository
}

func NewDarMeGustaComentarioUseCase(comentarioRepo domain.ComentarioRepository, userRepo domain.UserRepository) *DarMeGustaComentarioUseCase {
	return &DarMeGustaComentarioUseCase{
		comentarioRepo: comentarioRepo,
		userRepo:       userRepo,
	}
}

func (uc *DarMeGustaComentarioUseCase) Execute(usuarioID, comentarioID int64) *dto.ResponseThreads {
	user, _ := uc.userRepo.FindByID(usuarioID)
	if user == nil || !user.Exists() {
		return dto.NewResponseThreads(404, "Usuario no encontrado", nil)
	}

	comentario := uc.comentarioRepo.ObtenerComentario(comentarioID)
	if comentario == nil || !comentario.Existe() {
		return dto.NewResponseThreads(404, "Comentario no encontrado", nil)
	}

	ok := uc.comentarioRepo.DarMeGustaAComentario(usuarioID, comentarioID)
	if !ok {
		return dto.NewResponseThreads(500, "No se pudo registrar el me gusta", nil)
	}

	return dto.NewResponseThreads(200, "Me gusta registrado", nil)
}
