package middleware

type Response struct {
	RequestId string `json:"requestId"`
	Code      int    `json:"code"`
	Message   any    `json:"message,omitempty"`
	Data      any    `json:"data,omitempty"`
}

type DataResult struct {
	Total int64 `json:"total"`
	List  any   `json:"list"`
}
