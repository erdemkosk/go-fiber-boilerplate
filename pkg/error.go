package pkg

type Error struct {
	Status  int    `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

const (
	ErrExampleNotFound = 1001
)

var errorMessages = map[int]string{
	ErrExampleNotFound: "Example not found",
}

func NewError(status, code int) *Error {
	message, ok := errorMessages[code]
	if !ok {
		message = "Unknown error"
	}
	return &Error{
		Status:  status,
		Code:    code,
		Message: message,
	}
}

func ExampleNotFound() *Error {
	return NewError(404, ErrExampleNotFound)
}
