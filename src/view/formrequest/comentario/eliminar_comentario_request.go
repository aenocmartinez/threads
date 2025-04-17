package comentario

type EliminarComentarioRequest struct {
	UsuarioID    int64 `json:"usuario_id" binding:"required"`
	ComentarioID int64 `json:"comentario_id" binding:"required"`
}
