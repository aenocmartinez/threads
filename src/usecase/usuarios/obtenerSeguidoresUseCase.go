package usecase

import (
	"threads/src/domain"
	"threads/src/view/dto"
)

type ObtenerSeguidoresUseCase struct {
	userRepo domain.UserRepository
}

func NewObtenerSeguidoresUseCase(userRepo domain.UserRepository) *ObtenerSeguidoresUseCase {
	return &ObtenerSeguidoresUseCase{
		userRepo: userRepo,
	}
}

func (uc *ObtenerSeguidoresUseCase) Execute(usuarioID int64) *dto.ResponseThreads {
	seguidores := uc.userRepo.ObtenerUsuariosQueMeSiguen(usuarioID)
	if seguidores == nil {
		return dto.NewResponseThreads(200, "Sin seguidores", []dto.UserDTO{})
	}

	var resultado []dto.UserDTO
	for _, s := range *seguidores {
		if s.GetUserSeguidor() != nil {
			resultado = append(resultado, *s.GetUserSeguidor().ToDTO())
		}
	}

	return dto.NewResponseThreads(200, "Seguidores obtenidos correctamente", resultado)
}
