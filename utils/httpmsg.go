package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// NewError example
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

// Message example
type Message struct {
	Code    int         `json:"code" example:"1"`
	Message string      `json:"message" example:"ok"`
	Data    interface{} `json:"data"`
}

// NewMessage example
func NewMessageWithData(ctx *gin.Context, code int, message string, data interface{}) {
	Msg := Message{
		Code:    code,
		Message: message,
		Data:    data,
	}
	ctx.JSON(http.StatusOK, Msg)
}

// NewMessage example
func NewMessageWithError(ctx *gin.Context, err error) {
	var message string
	if err != nil {
		message = err.Error()
	}
	Msg := Message{
		Code:    0,
		Message: message,
	}
	ctx.JSON(http.StatusOK, Msg)
}

// NewMessage example
func NewMessage(ctx *gin.Context, code int, message string) {
	Msg := Message{
		Code:    code,
		Message: message,
	}
	ctx.JSON(http.StatusOK, Msg)
}
