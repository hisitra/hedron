package iotranslator

type Response struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func NewResponse(code uint32, message string, data string) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func SuccessResponse(message string) *Response {
	if message == "" {
		message = "Success"
	}
	return NewResponse(200, message, "{}")
}

func BadRequestResponse(message string) *Response {
	if message == "" {
		message = "Bad Request"
	}
	return NewResponse(400, message, "{}")
}

func InternalServerResponse(message string) *Response {
	if message == "" {
		message = "Internal Server Error"
	}
	return NewResponse(500, message, "{}")
}

func (r *Response) IsSuccessful() bool {
	return r.Code >= 200 && r.Code < 300
}

func (r *Response) Marshal() ([]byte, error) {
	return nil, nil
}
