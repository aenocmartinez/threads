package usecase

import (
	"threads/src/domain"
	"threads/src/view/dto"
)

type EliminarComentarioUseCase struct {
	comentarioRepo domain.ComentarioRepository
	userRepo       domain.UserRepository
}

func NewEliminarComentarioUseCase(comentarioRepo domain.ComentarioRepository, userRepo domain.UserRepository) *EliminarComentarioUseCase {
	return &EliminarComentarioUseCase{
		comentarioRepo: comentarioRepo,
		userRepo:       userRepo,
	}
}

func (uc *EliminarComentarioUseCase) Execute(usuarioID int64, comentarioID int64) *dto.ResponseThreads {
	user, _ := uc.userRepo.FindByID(usuarioID)
	if user == nil || !user.Exists() {
		return dto.NewResponseThreads(404, "Usuario no encontrado", nil)
	}

	comentario := uc.comentarioRepo.ObtenerComentario(comentarioID)
	if comentario == nil || !comentario.Existe() {
		return dto.NewResponseThreads(404, "Comentario no encontrado", nil)
	}

	if comentario.GetUser().GetID() != usuarioID {
		return dto.NewResponseThreads(403, "No tienes permiso para eliminar este comentario", nil)
	}

	ok := uc.comentarioRepo.EliminarComentario(comentarioID)
	if !ok {
		return dto.NewResponseThreads(500, "No se pudo eliminar el comentario", nil)
	}

	return dto.NewResponseThreads(200, "Comentario eliminado correctamente", nil)
}
