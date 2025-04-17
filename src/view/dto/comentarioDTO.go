package dto

import "time"

type ComentarioDTO struct {
	ID                int64          `json:"id,omitempty"`
	Usuario           *UserDTO       `json:"usuario,omitempty"`
	Contenido         string         `json:"contenido,omitempty"`
	ComentarioPadre   *ComentarioDTO `json:"comentario_padre,omitempty"`
	ComentarioPadreID *int64         `json:"comentario_padre_id,omitempty"`
	CreatedAt         time.Time      `json:"fecha_creacion,omitempty"`
	UpdatedAt         time.Time      `json:"fecha_actualizacion,omitempty"`
}
