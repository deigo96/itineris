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

type CreateEmployeeRequest struct {
	Name         string `json:"name" validate:"required"`
	NIP          string `json:"nip" validate:"required,numeric"`
	Password     string `json:"password" validate:"required"`
	Role         string `json:"role" validate:"required"`
	Position     string `json:"position" validate:"required"`
	Department   string `json:"department" validate:"required"`
	LeaveBalance int32  `json:"leave_balance" validate:"required"`
	IsPNS        bool   `json:"is_pns" validate:"required"`
}
