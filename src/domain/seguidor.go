package domain

import "time"

type Seguidor struct {
	userSeguidor *User
	userSeguido  *User
	fechaSigue   time.Time
}

func NewSeguidor() *Seguidor {
	return &Seguidor{}
}

func (s *Seguidor) SetUserSeguidor(userSeguidor *User) {
	s.userSeguidor = userSeguidor
}

func (s *Seguidor) GetUserSeguidor() *User {
	return s.userSeguidor
}

func (s *Seguidor) SetUserSeguido(userSeguido *User) {
	s.userSeguido = userSeguido
}

func (s *Seguidor) GetUserSeguido() *User {
	return s.userSeguido
}

func (s *Seguidor) SetFechaSigue(fechaSigue time.Time) {
	s.fechaSigue = fechaSigue
}

func (s *Seguidor) GetFechaSigue() time.Time {
	return s.fechaSigue
}
