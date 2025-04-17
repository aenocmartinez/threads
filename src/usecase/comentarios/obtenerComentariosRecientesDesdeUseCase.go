package usecase

import (
	"threads/src/domain"
	"threads/src/view/dto"
	"time"
)

type ObtenerComentariosRecientesDesdeUseCase struct {
	comentarioRepo domain.ComentarioRepository
}

func NewObtenerComentariosRecientesDesdeUseCase(comentarioRepo domain.ComentarioRepository) *ObtenerComentariosRecientesDesdeUseCase {
	return &ObtenerComentariosRecientesDesdeUseCase{
		comentarioRepo: comentarioRepo,
	}
}

func (uc *ObtenerComentariosRecientesDesdeUseCase) Execute(fecha time.Time) *dto.ResponseThreads {

	comentarios := uc.comentarioRepo.ObtenerComentariosRecientesDesde(fecha)

	return dto.NewResponseThreads(200, "Comentarios paginados desde fecha", comentarios)
}
