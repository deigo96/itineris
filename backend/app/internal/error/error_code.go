package error

import "errors"

var (
	// error 400
	ErrNotFound               = errors.New("not found")
	ErrBadRequest             = errors.New("bad request")
	ErrNipAlreadyUsed         = errors.New("nip already used")
	ErrIncorrectNipOrPassword = errors.New("nip or password incorrect")
	ErrUnauthorized           = errors.New("unauthorized")
	ErrLeaveBalance           = errors.New("leave balance not enough")
	ErrTimeLeaveRequest       = errors.New("time leave request not valid")
	ErrLeaveType              = errors.New("leave type not valid")

	// error 500
	ErrInternalServerError = errors.New("internal server error")
	ErrConflict            = errors.New("conflict")
)
