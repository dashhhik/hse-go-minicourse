package errs

import "fmt"

var (
	ErrInvalidRequest  = fmt.Errorf("invalid request")
	ErrAccountNotFound = fmt.Errorf("account not found")
)
