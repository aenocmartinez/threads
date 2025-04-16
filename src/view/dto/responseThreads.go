package dto

type ResponseThreads struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponseThreads(code int, message string, data interface{}) *ResponseThreads {
	return &ResponseThreads{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
