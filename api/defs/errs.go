package defs

type Err struct {
	error string `json:"error"`
	errorCode string `json:"error_code"`
}

type ErrorResponse struct {
	httpSC int
	error Err
}

var (
	errorRequestBodyParseFailed = ErrorResponse{httpSC: 400,error: Err{error:"Request body is not correct" , errorCode: "001"}}
	errorNotAuthUser = ErrorResponse{httpSC: 401,error: Err{error:"User authentication failed" , errorCode: "002"}}
)