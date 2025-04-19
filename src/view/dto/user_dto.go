package dto

type UserDTO struct {
	ID              int64  `json:"id,omitempty"`
	Name            string `json:"name"`
	Username        string `json:"username,omitempty"`
	Email           string `json:"email"`
	Avatar          string `json:"avatar"`
	Description     string `json:"description"`
	SessionToken    string `json:"token,omitempty"`
	TotalSeguidores int    `json:"total_seguidores"`
}
