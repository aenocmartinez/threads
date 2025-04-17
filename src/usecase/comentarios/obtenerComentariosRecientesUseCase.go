package usecase

import (
	"threads/src/domain"
	"threads/src/view/dto"
)

type ObtenerComentariosRecientesUseCase struct {
	comentarioRepo domain.ComentarioRepository
}

func NewObtenerComentariosRecientesUseCase(comentarioRepo domain.ComentarioRepository) *ObtenerComentariosRecientesUseCase {
	return &ObtenerComentariosRecientesUseCase{
		comentarioRepo: comentarioRepo,
	}
}

func (uc *ObtenerComentariosRecientesUseCase) Execute() *dto.ResponseThreads {

	comentarios := uc.comentarioRepo.ObtenerComentariosRecientes()

	return dto.NewResponseThreads(200, "Comentarios recientes", comentarios)
}
