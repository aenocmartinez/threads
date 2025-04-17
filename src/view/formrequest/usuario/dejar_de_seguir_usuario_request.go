package usuario

type DejarDeSeguirUsuarioRequest struct {
	SeguidorID int64 `json:"seguidor_id" binding:"required"`
	SeguidoID  int64 `json:"seguido_id" binding:"required"`
}
