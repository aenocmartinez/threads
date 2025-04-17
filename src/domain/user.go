package domain

import (
	"threads/src/view/dto"
)

type User struct {
	id           int64
	name         string
	username     string
	email        string
	phone        string
	password     string
	avatar       string
	description  string
	sessionToken string
	repository   UserRepository
}

func NewUser(repository UserRepository) *User {
	return &User{repository: repository}
}

func (u *User) SetID(id int64) {
	u.id = id
}

func (u *User) SetUsername(username string) {
	u.username = username
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) SetAvatar(avatar string) {
	u.avatar = avatar
}

func (u *User) SetDescription(description string) {
	u.description = description
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) SetSessionToken(sessionToken string) {
	u.sessionToken = sessionToken
}

func (u *User) SetPhone(phone string) {
	u.phone = phone
}

func (u *User) GetID() int64 {
	return u.id
}

func (u *User) GetUsername() string {
	return u.username
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetAvatar() string {
	return u.avatar
}

func (u *User) GetDescription() string {
	return u.description
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) GetSessionToken() string {
	return u.sessionToken
}

func (u *User) GetPhone() string {
	return u.phone
}

func (u *User) Exists() bool {
	return u.id > 0
}

func (u *User) Save() error {
	return u.repository.Save(u)
}

func (u *User) Update() error {
	return u.repository.Update(u)
}

func (u *User) Delete() error {
	return u.repository.Delete(u.id)
}

func (u *User) FindByID(id int64) (*User, error) {
	return u.repository.FindByID(id)
}

func (u *User) FindByEmail(email string) (*User, error) {
	return u.repository.FindByEmail(email)
}

func (u *User) FindByUsername(username string) (*User, error) {
	return u.repository.FindByUsername(username)
}

func (u *User) ObtenerUsuariosQueMeSiguen() *[]Seguidor {
	return u.repository.ObtenerUsuariosQueMeSiguen(u.id)
}

func (u *User) ObtenerUsuariosALosQueSigo() *[]Seguidor {
	return u.repository.ObtenerUsuariosQueSigo(u.id)
}

func (u *User) SeguirUsuario(usuarioASeguir *User) bool {
	return u.repository.SeguirUsuario(u.id, usuarioASeguir.GetID())
}

func (u *User) DejarDeSeguirUsuario(usuario_a_no_seguir *User) bool {
	return u.repository.DejarDeSeguirUsuario(u.id, usuario_a_no_seguir.GetID())
}

func (u *User) TotalDeSeguidores() int {
	return len(*u.ObtenerUsuariosQueMeSiguen())
}

func (u *User) ToDTO() *dto.UserDTO {
	return &dto.UserDTO{
		ID:              u.id,
		Name:            u.name,
		Avatar:          u.avatar,
		Username:        u.username,
		Email:           u.email,
		Description:     u.description,
		SessionToken:    u.sessionToken,
		TotalSeguidores: u.TotalDeSeguidores(),
	}
}
