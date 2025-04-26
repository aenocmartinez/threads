package usecase

import (
	"threads/src/domain"
	"threads/src/view/dto"
)

type ObtenerComentariosUsuarioUseCase struct {
	usuaRepo domain.UserRepository
}

func NewObtenerComentariosUsuarioUseCase(usuaRepo domain.UserRepository) *ObtenerComentariosUsuarioUseCase {
	return &ObtenerComentariosUsuarioUseCase{
		usuaRepo: usuaRepo,
	}
}

func (uc *ObtenerComentariosUsuarioUseCase) Execute(usuarioID int64) *dto.ResponseThreads {

	usuario, err := uc.usuaRepo.FindByID(usuarioID)
	if err != nil {
		return dto.NewResponseThreads(500, "Ha ocurrido un error al buscar un usuario", nil)
	}

	if !usuario.Exists() {
		return dto.NewResponseThreads(404, "Usuario no encontrado", nil)
	}

	comentarios := []dto.ComentarioDTO{}
	for _, comentario := range uc.usuaRepo.ObtenerComentariosRecientesUsuario(usuario.GetID()) {

		comentarios = append(comentarios, *comentario.MiComentarioToDTO())

	}

	return dto.NewResponseThreads(200, "Listado de comentarios", comentarios)
}
