package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

var db = make(map[string]string)
var router *gin.Engine

func init() {
  // mysql.Sync()
  router = gin.New()
  app := router.Group("/v1")
  v1.InitRoutes(app)
}

func noRouteHandler() gin.HandlerFunc{
  return func(c *gin.Context){
    var statusCode  int
    var message     string
    var data        interface{} = nil
    // var listError   []models.ErroNodel = nil
    var endPoint    string = c.Request.URL.String()
    var method      string = c.Request.Method
  }
}

func main() {
  http.ListenAndServe(":8081", router)
}
