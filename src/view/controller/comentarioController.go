package controller

import (
	"net/http"
	"strconv"

	"threads/src/infraestructure/di"
	usecase "threads/src/usecase/comentarios"
	"threads/src/view/formrequest/comentario"

	"time"

	"github.com/gin-gonic/gin"
)

func CrearComentario(c *gin.Context) {
	var request comentario.CrearComentarioRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	useCase := usecase.NewCrearComentarioUseCase(
		di.GetContainer().GetComentarioRepository(),
		di.GetContainer().GetUserRepository(),
	)

	response := useCase.Execute(request.UsuarioID, request.Contenido, request.ComentarioPadreID)
	c.JSON(http.StatusCreated, response)
}

func ActualizarComentario(c *gin.Context) {
	var request comentario.ActualizarComentarioRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	useCase := usecase.NewActualizarComentarioUseCase(
		di.GetContainer().GetComentarioRepository(),
		di.GetContainer().GetUserRepository(),
	)

	response := useCase.Execute(request.UsuarioID, request.ComentarioID, request.NuevoContenido)
	c.JSON(http.StatusOK, response)
}

func EliminarComentario(c *gin.Context) {

	var request comentario.EliminarComentarioRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	useCase := usecase.NewEliminarComentarioUseCase(
		di.GetContainer().GetComentarioRepository(),
		di.GetContainer().GetUserRepository(),
	)

	response := useCase.Execute(request.UsuarioID, request.ComentarioID)
	c.JSON(http.StatusOK, response)
}

func ResponderAComentario(c *gin.Context) {
	var request comentario.ResponderComentarioRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	useCase := usecase.NewResponderAComentarioUseCase(
		di.GetContainer().GetComentarioRepository(),
		di.GetContainer().GetUserRepository(),
	)

	response := useCase.Execute(request.UsuarioID, request.Contenido, request.ComentarioPadreID)
	c.JSON(http.StatusCreated, response)
}

func ObtenerConversacionDeComentario(c *gin.Context) {
	comentarioIDStr := c.Param("id")
	comentarioID, err := strconv.ParseInt(comentarioIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	useCase := usecase.NewObtenerConversacionDeComentarioUseCase(
		di.GetContainer().GetComentarioRepository(),
	)

	response := useCase.Execute(comentarioID)
	c.JSON(http.StatusOK, response)
}

func ObtenerComentariosRecientes(c *gin.Context) {
	comentarioRepo := di.GetContainer().GetComentarioRepository()
	useCase := usecase.NewObtenerComentariosRecientesUseCase(comentarioRepo)

	response := useCase.Execute()
	c.JSON(http.StatusOK, response)
}

func ObtenerComentariosRecientesDesde(c *gin.Context) {
	fechaStr := c.Query("antes_de")
	var fecha time.Time

	if fechaStr != "" {
		t, err := time.Parse("2006-01-02", fechaStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de fecha inválido (esperado YYYY-MM-DD)"})
			return
		}
		fecha = t
	}

	comentarioRepo := di.GetContainer().GetComentarioRepository()
	useCase := usecase.NewObtenerComentariosRecientesDesdeUseCase(comentarioRepo)

	response := useCase.Execute(fecha)
	c.JSON(http.StatusOK, response)
}

func DarMeGustaComentario(c *gin.Context) {
	comentarioIDStr := c.Param("id")
	comentarioID, err := strconv.ParseInt(comentarioIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de comentario inválido"})
		return
	}

	var request comentario.MeGustaComentarioRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	useCase := usecase.NewDarMeGustaComentarioUseCase(
		di.GetContainer().GetComentarioRepository(),
		di.GetContainer().GetUserRepository(),
	)

	response := useCase.Execute(request.UsuarioID, comentarioID)
	c.JSON(http.StatusOK, response)
}

func QuitarMeGustaComentario(c *gin.Context) {
	comentarioIDStr := c.Param("id")
	comentarioID, err := strconv.ParseInt(comentarioIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de comentario inválido"})
		return
	}

	var request comentario.MeGustaComentarioRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	useCase := usecase.NewQuitarMeGustaComentarioUseCase(
		di.GetContainer().GetComentarioRepository(),
		di.GetContainer().GetUserRepository(),
	)

	response := useCase.Execute(request.UsuarioID, comentarioID)
	c.JSON(http.StatusOK, response)
}

func ObtenerUsuariosQueDieronMeGusta(c *gin.Context) {
	comentarioIDStr := c.Param("id")
	comentarioID, err := strconv.ParseInt(comentarioIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	useCase := usecase.NewObtenerUsuariosQueDieronMeGustaUseCase(
		di.GetContainer().GetComentarioRepository(),
	)

	response := useCase.Execute(comentarioID)
	c.JSON(http.StatusOK, response)
}
