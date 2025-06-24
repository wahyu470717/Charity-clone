// response/response.go
package response

import "net/http"

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type APIResponse struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta"`
}

func SuccessResponse(data interface{}) APIResponse {
	return APIResponse{
		Data: data,
		Meta: Meta{
			Code:    http.StatusOK,
			Message: "Success",
			Status:  "success",
		},
	}
}

func ErrorResponse(message string) Meta {
	return Meta{
		Code:    http.StatusBadRequest,
		Message: message,
		Status:  "error",
	}
}

func ErrorResponseWithCode(code int, message string) Meta {
	return Meta{
		Code:    code,
		Message: message,
		Status:  "error",
	}
}
