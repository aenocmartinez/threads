package comentario

type CrearComentarioRequest struct {
	UsuarioID         int64  `json:"usuario_id" binding:"required"`
	Contenido         string `json:"contenido" binding:"required"`
	ComentarioPadreID *int64 `json:"comentario_padre_id,omitempty"`
}
