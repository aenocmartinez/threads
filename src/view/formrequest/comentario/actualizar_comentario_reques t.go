package comentario

type ActualizarComentarioRequest struct {
	UsuarioID      int64  `json:"usuario_id" binding:"required"`
	ComentarioID   int64  `json:"comentario_id" binding:"required"`
	NuevoContenido string `json:"contenido" binding:"required"`
}
