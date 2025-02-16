package utils

import (
	"github.com/labstack/echo/v4"
)

type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func JSONResponse(ctx echo.Context, code int, message string, data interface{}, err error) error {
	response := APIResponse{
		Code:    code,
		Message: message,
	}

	if err != nil {
		response.Error = err.Error()
	} else {
		response.Data = data
	}

	return ctx.JSON(code, response)
}
