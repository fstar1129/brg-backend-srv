package eth

// Error ...
type Error struct {
	err       error
	retryable bool
}

// NewError ...
func NewError(err error, retryable bool) *Error {
	return &Error{
		err:       err,
		retryable: retryable,
	}
}
