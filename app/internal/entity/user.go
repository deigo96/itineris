package entity

type UserResponse struct {
	Email     string `json:"email"`
	IsActive  bool   `json:"is_active"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"created_by"`
	UpdatedAt string `json:"updated_at"`
	UpdatedBy string `json:"updated_by"`
}
