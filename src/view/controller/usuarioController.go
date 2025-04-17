package controller

import (
	"net/http"
	"strconv"
	"threads/src/infraestructure/di"
	usecase "threads/src/usecase/usuarios"
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
