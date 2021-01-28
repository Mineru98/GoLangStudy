package cors

import (
	"github.com/gin-gonic/gin"
)

func New(c *gin.Context) {
    c.Header("Access-Control-Allow-Credentials", "True")
    c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE")
    c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Origin, Accept, Cache-Control, X-Requested-With")
    c.Next()
}