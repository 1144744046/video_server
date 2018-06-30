package main

type ApiBody struct {
	Url string `json:"url"`
	Method string `json:"method"`
	ReqBody string `json:"req_body"`
}

type Error struct {
	Err string `json:"err"`
	ErrCode string `json:"errcode"`
}

var (
	ErrorBadRequest=Error{"api bad request","001"}
	ErrorRequestBodyParseFail=Error{"request body not current","002"}
	ErrorResponseBodyParseFail=Error{"response body not current","003"}

	)