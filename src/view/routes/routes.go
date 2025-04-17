package routes

import (
	"threads/src/infraestructure/middleware"
	"threads/src/view/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Rutas públicas
	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)
	r.GET("/check-db", controller.CheckDBConnection)

	// Rutas protegidas
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/logout", controller.Logout)

		// Comentarios
		protected.GET("/comentarios", controller.ObtenerComentariosRecientes)
		protected.GET("/comentarios/feed", controller.ObtenerComentariosRecientesDesde)
		protected.GET("/comentarios/:id/conversacion", controller.ObtenerConversacionDeComentario)

		protected.POST("/comentarios", controller.CrearComentario)
		protected.PUT("/comentarios", controller.ActualizarComentario)
		protected.DELETE("/comentarios", controller.EliminarComentario)
		protected.POST("/comentarios/responder", controller.ResponderAComentario)
		protected.POST("/comentarios/:id/like", controller.DarMeGustaComentario)
		protected.DELETE("/comentarios/:id/like", controller.QuitarMeGustaComentario)
		protected.GET("/comentarios/:id/likes", controller.ObtenerUsuariosQueDieronMeGusta)

		// Usuarios
		protected.POST("/usuarios/seguir", controller.SeguirUsuario)
		protected.POST("/usuarios/dejar-de-seguir", controller.DejarDeSeguirUsuario)
		protected.GET("/usuarios/:id/seguidores", controller.ObtenerSeguidores)
		protected.GET("/usuarios/:id/seguidos", controller.ObtenerSeguidos)
		protected.PUT("/usuarios/perfil", controller.EditarPerfil)
		protected.POST("/usuarios/avatar", controller.SubirAvatar)
		protected.GET("/usuarios/:id", controller.ConsultarUsuario)
	}
}
