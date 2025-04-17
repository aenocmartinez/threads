package usecase

import (
	"threads/src/domain"
	"threads/src/view/dto"
)

type ObtenerUsuariosQueDieronMeGustaUseCase struct {
	comentarioRepo domain.ComentarioRepository
}

func NewObtenerUsuariosQueDieronMeGustaUseCase(comentarioRepo domain.ComentarioRepository) *ObtenerUsuariosQueDieronMeGustaUseCase {
	return &ObtenerUsuariosQueDieronMeGustaUseCase{comentarioRepo: comentarioRepo}
}

func (uc *ObtenerUsuariosQueDieronMeGustaUseCase) Execute(comentarioID int64) *dto.ResponseThreads {
	usuarios := uc.comentarioRepo.ObtenerUsuariosQueDieronMeGusta(comentarioID)

	resultado := []dto.UsuarioMeGustoDTO{}
	for _, u := range usuarios {
		resultado = append(resultado, dto.UsuarioMeGustoDTO{
			ID:       u.GetID(),
			Name:     u.GetName(),
			Username: u.GetUsername(),
			Email:    u.GetEmail(),
			Avatar:   u.GetAvatar(),
		})
	}

	return dto.NewResponseThreads(200, "Usuarios que dieron me gusta", resultado)
}
