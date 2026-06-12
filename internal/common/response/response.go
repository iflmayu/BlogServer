package response

import (
	"BlogServer/pkg/utils/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code Code   `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func OkWithData(data any, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: SuccessCode,
		Msg:  SuccessCode.String(),
		Data: data,
	})
}

func OkWithMsg(msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: SuccessCode,
		Msg:  msg,
		Data: map[string]any{},
	})
}

func OkWithList(list any, count int, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: SuccessCode,
		Data: map[string]any{
			"list":  list,
			"count": count,
		},
		Msg: "成功",
	})
}

func FailWithMsg(msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: FailValidCode,
		Msg:  msg,
		Data: map[string]any{},
	})
}

func FailWithCode(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: FailServiceCode,
		Msg:  FailServiceCode.String(),
		Data: map[string]any{},
	})
}

func FailWithData(data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: FailServiceCode,
		Msg:  msg,
		Data: data,
	})
}

func FailWithError(err error, c *gin.Context) {
	msg := validator.ValidateErr(err)
	data := validator.ValidateFieldErrors(err)
	FailWithData(data, msg, c)
}
