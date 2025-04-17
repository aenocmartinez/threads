package dto

type UserDTO struct {
	ID              int64  `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Username        string `json:"username,omitempty"`
	Email           string `json:"email,omitempty"`
	Avatar          string `json:"avatar,omitempty"`
	Description     string `json:"description,omitempty"`
	SessionToken    string `json:"token,omitempty"`
	TotalSeguidores int    `json:"total_seguidores"`
}
