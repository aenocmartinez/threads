package domain

import (
	"threads/src/view/dto"
	"time"
)

type Comentario struct {
	id                   int64
	user                 *User
	contenido            string
	comentarioPadre      *Comentario
	createdAt            time.Time
	updatedAt            time.Time
	userRepository       UserRepository
	comentarioRepository ComentarioRepository
}

func NewComentario(userRepository UserRepository, comentarioRepository ComentarioRepository) *Comentario {
	return &Comentario{
		userRepository:       userRepository,
		comentarioRepository: comentarioRepository,
	}
}

func (c *Comentario) SetID(id int64) {
	c.id = id
}

func (c *Comentario) GetID() int64 {
	return c.id
}

func (c *Comentario) SetUser(user *User) {
	c.user = user
}

func (c *Comentario) GetUser() *User {
	return c.user
}

func (c *Comentario) SetContenido(contenido string) {
	c.contenido = contenido
}

func (c *Comentario) GetContenido() string {
	return c.contenido
}

func (c *Comentario) SetComentarioPadre(comentarioPadre *Comentario) {
	c.comentarioPadre = comentarioPadre
}

func (c *Comentario) GetComentarioPadre() *Comentario {
	return c.comentarioPadre
}

func (c *Comentario) SetCreatedAt(createdAt time.Time) {
	c.createdAt = createdAt
}

func (c *Comentario) GetCreatedAt() time.Time {
	return c.createdAt
}

func (c *Comentario) SetUpdatedAt(updatedAt time.Time) {
	c.updatedAt = updatedAt
}

func (c *Comentario) GetUpdatedAt() time.Time {
	return c.updatedAt
}

func (c *Comentario) Existe() bool {
	return c.id > 0
}

func (c *Comentario) ToDTO() *dto.ComentarioDTO {
	var comentarioPadre *dto.ComentarioDTO
	if c.comentarioPadre != nil {
		comentarioPadre = c.comentarioPadre.ToDTO()
	}

	return &dto.ComentarioDTO{
		ID:              c.GetID(),
		Contenido:       c.GetContenido(),
		Usuario:         c.user.ToDTO(),
		ComentarioPadre: comentarioPadre,
		CreatedAt:       c.GetCreatedAt(),
	}
}
