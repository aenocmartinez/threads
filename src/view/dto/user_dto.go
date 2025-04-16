package dto

type UserDTO struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Avatar       string `json:"avatar"`
	Description  string `json:"description"`
	SessionToken string `json:"token"`
}
