package errors

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewAppError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func (e *AppError) Error() string {
	return e.Message
}

func New(code int, msg string) error {
	return &AppError{
		Code:    code,
		Message: msg,
	}
}

func Wrap(code int, err error) error {
	return &AppError{
		Code:    code,
		Message: err.Error(),
	}
}
