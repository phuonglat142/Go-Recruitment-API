package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestIdKey = "requestId"

func RequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.GetHeader("x-request-id")
		if id == "" {
			id = uuid.NewString()
		}
		ctx.Set(RequestIdKey, id)
		ctx.Header("x-request-id", id)
		ctx.Next()
	}
}

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()

		log.Printf("[%s] %s %s | %d | %v",
			ctx.GetString(RequestIdKey),
			ctx.Request.Method,
			ctx.Request.URL.Path,
			ctx.Writer.Status(),
			time.Since(start),
		)
	}
}
