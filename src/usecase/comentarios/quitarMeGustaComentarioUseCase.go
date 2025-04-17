package usecase

import (
	"threads/src/domain"
	"threads/src/view/dto"
)

type QuitarMeGustaComentarioUseCase struct {
	comentarioRepo domain.ComentarioRepository
	userRepo       domain.UserRepository
}

func NewQuitarMeGustaComentarioUseCase(comentarioRepo domain.ComentarioRepository, userRepo domain.UserRepository) *QuitarMeGustaComentarioUseCase {
	return &QuitarMeGustaComentarioUseCase{
		comentarioRepo: comentarioRepo,
		userRepo:       userRepo,
	}
}

func (uc *QuitarMeGustaComentarioUseCase) Execute(usuarioID, comentarioID int64) *dto.ResponseThreads {
	user, _ := uc.userRepo.FindByID(usuarioID)
	if user == nil || !user.Exists() {
		return dto.NewResponseThreads(404, "Usuario no encontrado", nil)
	}

	comentario := uc.comentarioRepo.ObtenerComentario(comentarioID)
	if comentario == nil || !comentario.Existe() {
		return dto.NewResponseThreads(404, "Comentario no encontrado", nil)
	}

	ok := uc.comentarioRepo.QuitarMeGustaAComentario(usuarioID, comentarioID)
	if !ok {
		return dto.NewResponseThreads(500, "No se pudo quitar el me gusta", nil)
	}

	return dto.NewResponseThreads(200, "Me gusta eliminado", nil)
}
