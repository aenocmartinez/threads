package domain

type MeGustaComentario struct {
	user       *User
	comentario *Comentario
}

func NewMeGustaComentario(user *User, comentario *Comentario) *MeGustaComentario {
	return &MeGustaComentario{
		user:       user,
		comentario: comentario,
	}
}

func (mc *MeGustaComentario) GetUser() *User {
	return mc.user
}

func (mc *MeGustaComentario) GetComentario() *Comentario {
	return mc.comentario
}
