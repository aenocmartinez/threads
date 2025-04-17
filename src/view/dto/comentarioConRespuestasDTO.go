package dto

type ComentarioConRespuestasDTO struct {
	Comentario ComentarioDTO                 `json:"comentario"`
	Respuestas []*ComentarioConRespuestasDTO `json:"respuestas"`
}
