package usecase

import (
	"threads/src/domain"
	"threads/src/view/dto"
)

type ResponderAComentarioUseCase struct {
	comentarioRepo domain.ComentarioRepository
	userRepo       domain.UserRepository
}

func NewResponderAComentarioUseCase(comentarioRepo domain.ComentarioRepository, userRepo domain.UserRepository) *ResponderAComentarioUseCase {
	return &ResponderAComentarioUseCase{
		comentarioRepo: comentarioRepo,
		userRepo:       userRepo,
	}
}

func (uc *ResponderAComentarioUseCase) Execute(usuarioID int64, contenido string, comentarioPadreID int64) *dto.ResponseThreads {
	user, _ := uc.userRepo.FindByID(usuarioID)
	if user == nil || !user.Exists() {
		return dto.NewResponseThreads(404, "Usuario no encontrado", nil)
	}

	comentarioPadre := uc.comentarioRepo.ObtenerComentario(comentarioPadreID)
	if comentarioPadre == nil || !comentarioPadre.Existe() {
		return dto.NewResponseThreads(404, "Comentario padre no encontrado", nil)
	}

	respuesta := domain.NewComentario(uc.userRepo, uc.comentarioRepo)
	respuesta.SetUser(user)
	respuesta.SetContenido(contenido)
	respuesta.SetComentarioPadre(comentarioPadre)

	ok := uc.comentarioRepo.ResponderAComentario(comentarioPadreID, respuesta)
	if !ok {
		return dto.NewResponseThreads(500, "No se pudo registrar la respuesta", nil)
	}

	return dto.NewResponseThreads(201, "Comentario respondido correctamente", respuesta.ToDTO())
}
