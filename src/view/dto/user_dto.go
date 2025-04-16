package dto

type UserDTO struct {
	ID           int64  `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	SessionToken string `json:"token"`
}
