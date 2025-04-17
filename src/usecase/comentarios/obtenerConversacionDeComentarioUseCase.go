package usecase

import (
	"threads/src/domain"
	"threads/src/view/dto"
)

type ObtenerConversacionDeComentarioUseCase struct {
	comentarioRepo domain.ComentarioRepository
}

func NewObtenerConversacionDeComentarioUseCase(comentarioRepo domain.ComentarioRepository) *ObtenerConversacionDeComentarioUseCase {
	return &ObtenerConversacionDeComentarioUseCase{
		comentarioRepo: comentarioRepo,
	}
}

func (uc *ObtenerConversacionDeComentarioUseCase) Execute(comentarioID int64) *dto.ResponseThreads {
	comentario := uc.comentarioRepo.ObtenerComentario(comentarioID)
	if comentario == nil || !comentario.Existe() {
		return dto.NewResponseThreads(404, "Comentario no encontrado", nil)
	}

	for comentario.GetComentarioPadre() != nil {
		comentario = comentario.GetComentarioPadre()
	}

	conversacion := uc.comentarioRepo.ObtenerConversacion(comentario.GetID())

	return dto.NewResponseThreads(200, "Conversación encontrada", conversacion)
}
