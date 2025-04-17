package comentario

type MeGustaComentarioRequest struct {
	UsuarioID int64 `json:"usuario_id" binding:"required"`
}
