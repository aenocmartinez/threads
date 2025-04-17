package usuario

type EditarPerfilRequest struct {
	ID   int64  `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	// Username    string `json:"username" binding:"required"`
	// Email       string `json:"email" binding:"required"`
	Phone       string `json:"phone"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
}
