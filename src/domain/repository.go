package domain

import (
	"threads/src/view/dto"
	"time"
)

type UserRepository interface {
	FindByID(id int64) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByUsername(username string) (*User, error)
	FindUserLogin(login string) (*User, error)
	ExistsUsername(username string) (bool, error)
	Save(user *User) error
	Update(user *User) error
	Delete(id int64) error
	ObtenerUsuariosQueMeSiguen(userID int64) []Seguidor
	ObtenerUsuariosQueSigo(userID int64) []Seguidor
	SeguirUsuario(usuarioSeguidorID, usuarioSeguidoID int64) bool
	DejarDeSeguirUsuario(usuarioSeguidorID, usuarioSeguidoID int64) bool
}

type ComentarioRepository interface {
	CrearComentario(comentario *Comentario) bool
	ResponderAComentario(comentarioPadreID int64, respuesta *Comentario) bool
	ActualizarComentario(comentario *Comentario) bool
	EliminarComentario(comentarioID int64) bool
	ObtenerComentario(id int64) *Comentario
	ObtenerConversacion(comentarioID int64) dto.ComentarioConRespuestasDTO
	ObtenerComentariosRecientes() []dto.ComentarioConRespuestasDTO
	ObtenerComentariosRecientesDesde(fechaUltimo time.Time) []dto.ComentarioConRespuestasDTO
	DarMeGustaAComentario(usuarioID, comentarioID int64) bool
	QuitarMeGustaAComentario(usuarioID, comentarioID int64) bool
	ObtenerUsuariosQueDieronMeGusta(comentarioID int64) []User
}
