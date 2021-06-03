//routes.go
package mainrouter

import (
  "github.com/gin-gonic/gin"
  "../Controller"
)


func SetupRoute(router *gin.Engine){
    router.GET("/", controller.ShowIndexPage)
}

