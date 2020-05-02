package iotranslator

import "github.com/hisitra/hedron/src/comcn"

func NewResponse(code uint32, message string, data []byte) *comcn.Output {
	return &comcn.Output{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func EarlierRequestFoundResponse() *comcn.Output {
	return NewResponse(408, "earlier request found", nil)
}

func SuccessResponse(message string) *comcn.Output {
	if message == "" {
		message = "Success"
	}
	return NewResponse(200, message, nil)
}

func NotFoundResponse(message string) *comcn.Output {
	if message == "" {
		message = "Not Found"
	}
	return NewResponse(404, message, nil)
}

func BadRequestResponse(message string) *comcn.Output {
	if message == "" {
		message = "Bad Request"
	}
	return NewResponse(400, message, nil)
}

func UnauthorizedResponse(message string) *comcn.Output {
	if message == "" {
		message = "Unauthorized"
	}
	return NewResponse(401, message, nil)
}

func TimeoutResponse(message string) *comcn.Output {
	if message == "" {
		message = "Request Timed Out."
	}
	return NewResponse(408, message, nil)
}

func InternalServerErrorResponse(message string) *comcn.Output {
	if message == "" {
		message = "Internal Server Error"
	}
	return NewResponse(500, message, nil)
}
