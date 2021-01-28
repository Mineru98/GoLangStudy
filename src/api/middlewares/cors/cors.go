package cors

import (
	"github.com/gin-gonic/gin"
)

func CorsMiddleware(c *gin.Context) {
    c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
    c.Header("Access-Control-Allow-Credentials", "True")
    c.Header("Access-Control-Allow-Origin", "http://localhost:8081")
    c.Header("Access-Control-Allow-Methods", "GET, POST, DELETE")
    c.Next()
}