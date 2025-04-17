package controller

import (
	"net/http"
	"strconv"
	"threads/src/infraestructure/di"
	usecase "threads/src/usecase/usuarios"
	"threads/src/view/dto"
	"threads/src/view/formrequest/usuario"

	"github.com/gin-gonic/gin"
)

func SeguirUsuario(c *gin.Context) {

	var request usuario.SeguirUsuarioRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	userRepo := di.GetContainer().GetUserRepository()
	useCase := usecase.NewSeguirUsuarioUseCase(userRepo)

	response := useCase.Execute(request.SeguidorID, request.SeguidoID)
	c.JSON(http.StatusOK, response)
}

func DejarDeSeguirUsuario(c *gin.Context) {
	var request usuario.DejarDeSeguirUsuarioRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	userRepo := di.GetContainer().GetUserRepository()
	useCase := usecase.NewDejarDeSeguirUsuarioUseCase(userRepo)

	response := useCase.Execute(request.SeguidorID, request.SeguidoID)
	c.JSON(http.StatusOK, response)
}

func ObtenerSeguidores(c *gin.Context) {
	usuarioIDStr := c.Param("id")
	usuarioID, err := strconv.ParseInt(usuarioIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	useCase := usecase.NewObtenerSeguidoresUseCase(
		di.GetContainer().GetUserRepository(),
	)

	response := useCase.Execute(usuarioID)
	c.JSON(http.StatusOK, response)
}

func ObtenerSeguidos(c *gin.Context) {
	usuarioIDStr := c.Param("id")
	usuarioID, err := strconv.ParseInt(usuarioIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	useCase := usecase.NewObtenerSeguidosUseCase(
		di.GetContainer().GetUserRepository(),
	)

	response := useCase.Execute(usuarioID)
	c.JSON(http.StatusOK, response)
}

func EditarPerfil(c *gin.Context) {
	var req usuario.EditarPerfilRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	useCase := usecase.NewEditarPerfilUseCase(di.GetContainer().GetUserRepository())
	dtoInput := dto.EditarPerfilDTO{
		ID:   req.ID,
		Name: req.Name,
		// Username:    req.Username,
		// Email:       req.Email,
		Phone:       req.Phone,
		Avatar:      req.Avatar,
		Description: req.Description,
	}

	response := useCase.Execute(dtoInput)
	c.JSON(http.StatusOK, response)
}

func SubirAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Archivo avatar no encontrado"})
		return
	}

	useCase := usecase.NewSubirAvatarUseCase()
	path, err := useCase.Execute(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"path": path})
}
