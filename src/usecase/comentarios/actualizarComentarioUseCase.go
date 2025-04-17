package usecase

import (
	"threads/src/domain"
	"threads/src/view/dto"
)

type ActualizarComentarioUseCase struct {
	comentarioRepo domain.ComentarioRepository
	userRepo       domain.UserRepository
}

func NewActualizarComentarioUseCase(comentarioRepo domain.ComentarioRepository, userRepo domain.UserRepository) *ActualizarComentarioUseCase {
	return &ActualizarComentarioUseCase{
		comentarioRepo: comentarioRepo,
		userRepo:       userRepo,
	}
}

func (uc *ActualizarComentarioUseCase) Execute(usuarioID int64, comentarioID int64, nuevoContenido string) *dto.ResponseThreads {
	user, _ := uc.userRepo.FindByID(usuarioID)
	if user == nil || !user.Exists() {
		return dto.NewResponseThreads(404, "Usuario no encontrado", nil)
	}

	comentario := uc.comentarioRepo.ObtenerComentario(comentarioID)
	if comentario == nil || !comentario.Existe() {
		return dto.NewResponseThreads(404, "Comentario no encontrado", nil)
	}

	if comentario.GetUser().GetID() != usuarioID {
		return dto.NewResponseThreads(403, "No tienes permiso para editar este comentario", nil)
	}

	comentario.SetContenido(nuevoContenido)

	ok := uc.comentarioRepo.ActualizarComentario(comentario)
	if !ok {
		return dto.NewResponseThreads(500, "No se pudo actualizar el comentario", nil)
	}

	return dto.NewResponseThreads(200, "Comentario actualizado correctamente", comentario.ToDTO())
}
