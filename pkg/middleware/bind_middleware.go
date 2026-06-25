package middleware

import (
	"github.com/gin-gonic/gin"

	"BlogServer/internal/common/response"
)

func BindJSON[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req T
		if err := c.ShouldBindJSON(&req); err != nil {
			response.FailWithError(err, c)
			c.Abort()
			return
		}
		c.Set("request", req)
		c.Next()
	}
}

func BindQuery[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req T
		if err := c.ShouldBindQuery(&req); err != nil {
			response.FailWithError(err, c)
			c.Abort()
			return
		}
		c.Set("request", req)
		c.Next()
	}
}

func BindUri[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req T
		if err := c.ShouldBindUri(&req); err != nil {
			response.FailWithError(err, c)
			c.Abort()
			return
		}
		c.Set("request", req)
		c.Next()
	}
}

func GetRequest[T any](c *gin.Context) T {
	req, _ := c.Get("request")
	return req.(T)
}
