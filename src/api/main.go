package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "api/routers/v1"
  "api/middlewares/cors"
)

var db = make(map[string]string)
var router *gin.Engine

func init() {
  // mysql.Sync()
  router = gin.New()
  router.Use(cors.New) // Cors 정책 설정
  app := router.Group("/v1")
  v1.InitRoutes(app)
}

func noRouteHandler() gin.HandlerFunc{
  return func(c *gin.Context){
    // var statusCode  int
    // var message     string
    // var data        interface{} = nil
    // // var listError   []models.ErroNodel = nil
    // var endPoint    string = c.Request.URL.String()
    // var method      string = c.Request.Method
  }
}

func main() {
  http.ListenAndServe(":8081", router)
}
