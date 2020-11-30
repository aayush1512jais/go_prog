package apperrors

type ErrorModel struct {
	Message string
	Error   error
	Code    int
}

// type ErrorHandler interface {
// 	Response(w http.ResponseWriter, r *http.Request)
// }

// func NewErrorHandler() ErrorHandler {
// 	return &ErrorModel{}
// }

// func

// // func (err *ErrorModel) Response(w http.ResponseWriter, r *http.Request) {

// // 	json.NewEncoder(w).Encode(&err)
// // 	return
// // }
