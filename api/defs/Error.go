package defs

type Error struct {
	ErrString string `json:"err_string"`
	ErrCode string `json:"err_code"`
}

type ErrorResponse struct {
	HttpRC int
	Err Error
}

var (
	ErrorRequestBodyParseFailed=ErrorResponse{400,Error{"Request body parse fail","001"}}
	ErrorUserNotAuth=ErrorResponse{401,Error{"User auth fail","002"}}
	ErrorDBError=ErrorResponse{500,Error{"Db ops fail","003"}}
	ErrorInternalFaults=ErrorResponse{500,Error{"Internal service error","004"}}

)