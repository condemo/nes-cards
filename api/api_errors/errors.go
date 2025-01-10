package apiErrors

type ApiError struct {
	Err    error
	Msg    string
	Status int
}

func (e ApiError) Error() string {
	return e.Err.Error()
}

func NewApiError(err error, msg string, status int) ApiError {
	return ApiError{Err: err, Msg: msg, Status: status}
}
