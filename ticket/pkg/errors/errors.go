package errors

import "github.com/pkg/errors"

var (
	ErrInvalidID           = errors.New("Invalid uuid")
	ErrInternalServerError = errors.New("Internal server error")
	ErrTicketNotFound      = errors.New("Ticket not found")
)
