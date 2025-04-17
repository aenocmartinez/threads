package comentario

type ResponderComentarioRequest struct {
	UsuarioID         int64  `json:"usuario_id" binding:"required"`
	ComentarioPadreID int64  `json:"comentario_padre_id" binding:"required"`
	Contenido         string `json:"contenido" binding:"required"`
}
