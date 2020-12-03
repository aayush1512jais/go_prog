package apperrors

type ErrorModel struct {
	Message string
	Error   error
	Code    int
}
