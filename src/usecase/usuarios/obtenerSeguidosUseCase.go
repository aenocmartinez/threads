package usecase

import (
	"threads/src/domain"
	"threads/src/view/dto"
)

type ObtenerSeguidosUseCase struct {
	userRepo domain.UserRepository
}

func NewObtenerSeguidosUseCase(userRepo domain.UserRepository) *ObtenerSeguidosUseCase {
	return &ObtenerSeguidosUseCase{
		userRepo: userRepo,
	}
}

func (uc *ObtenerSeguidosUseCase) Execute(usuarioID int64) *dto.ResponseThreads {
	seguidos := uc.userRepo.ObtenerUsuariosQueSigo(usuarioID)
	if seguidos == nil {
		return dto.NewResponseThreads(200, "Sin usuarios seguidos", []dto.UserDTO{})
	}

	resultado := []dto.UserDTO{}
	for _, s := range seguidos {
		if s.GetUserSeguido() != nil {
			resultado = append(resultado, *s.GetUserSeguido().ToDTO())
		}
	}

	return dto.NewResponseThreads(200, "Usuarios seguidos obtenidos correctamente", resultado)
}
