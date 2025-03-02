package model

type EmployeeResponse struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	NIP                 string `json:"nip"`
	LeaveBalance        int32  `json:"leave_balance"`
	Role                string `json:"role"`
	TotalPendingRequest int    `json:"total_pending_request"`
	Position            string `json:"position"`
	Department          string `json:"department"`
	CreatedAt           string `json:"created_at"`
	CreatedBy           string `json:"created_by"`
	UpdatedAt           string `json:"updated_at"`
	UpdatedBy           string `json:"updated_by"`
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
