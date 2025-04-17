package usecase

import (
	"threads/src/domain"
	"threads/src/view/dto"
)

type CrearComentarioUseCase struct {
	comentarioRepo domain.ComentarioRepository
	userRepo       domain.UserRepository
}

func NewCrearComentarioUseCase(comentarioRepo domain.ComentarioRepository, userRepo domain.UserRepository) *CrearComentarioUseCase {
	return &CrearComentarioUseCase{
		comentarioRepo: comentarioRepo,
		userRepo:       userRepo,
	}
}

func (uc *CrearComentarioUseCase) Execute(usuarioID int64, contenido string, comentarioPadreID *int64) *dto.ResponseThreads {
	user, _ := uc.userRepo.FindByID(usuarioID)
	if user == nil || !user.Exists() {
		return dto.NewResponseThreads(404, "Usuario no encontrado", nil)
	}

	comentario := domain.NewComentario(uc.userRepo, uc.comentarioRepo)
	comentario.SetUser(user)
	comentario.SetContenido(contenido)

	if comentarioPadreID != nil {
		padre := uc.comentarioRepo.ObtenerComentario(*comentarioPadreID)
		if padre != nil && padre.Existe() {
			comentario.SetComentarioPadre(padre)
		}
	}

	ok := uc.comentarioRepo.CrearComentario(comentario)
	if !ok {
		return dto.NewResponseThreads(500, "No se pudo crear el comentario", nil)
	}

	return dto.NewResponseThreads(201, "Comentario creado correctamente", comentario.ToDTO())
}
