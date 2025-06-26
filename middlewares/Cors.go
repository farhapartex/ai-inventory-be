package middlewares

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
		if allowedOrigins == "" {
			allowedOrigins = "*"
		}

		origin := ctx.Request.Header.Get("Origin")
		if allowedOrigins == "*" {
			ctx.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			for _, allowedOrigin := range strings.Split(allowedOrigins, ",") {
				if strings.TrimSpace(allowedOrigin) == origin {
					ctx.Writer.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
					break
				}

			}
		}

		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204) // No content
			return
		}

		ctx.Next()
	}
}
