package dto

type MeGustaComentarioDTO struct {
	User       UserDTO       `json:"usuario"`
	Comentario ComentarioDTO `json:"comentario"`
}
