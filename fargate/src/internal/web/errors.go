package web

type ErrorCode string

const (
	ValidationErrCode     ErrorCode = "Validation"
	ConflictErrorCode     ErrorCode = "Conflict"
	NotFoundErrCode       ErrorCode = "NotFound"
	InternalServerErrCode ErrorCode = "Internal"
)

type Error struct {
	Err       error     `json:"-"`
	ErrorType ErrorCode `json:"errorType,omitempty"`
	Message   string    `json:"message,omitempty"`
	Info      string    `json:"info,omitempty"`
}

func (e *Error) Error() string {
	return e.Err.Error()
}
