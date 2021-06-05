//routes.go
package mainrouters

import (
  "github.com/gin-gonic/gin"
  mc "../maincontrollers"
)


func SetupRoute(router *gin.Engine){
    router.GET("/", mc.ShowIndexPage)
    router.GET("/users", mc.GetUsers)
    router.GET("/user/:id", mc.UserController.GetUserByID)
    router.POST("/user/", mc.UserController.GetUserByID)
    router.PUT("/user/:id", mc.UserController.GetUserByID)
    router.DELETE("/user/:id", mc.UserController.GetUserByID)
    router.DELETE("/user/:id/confirm", mc.UserController.GetUserByID)
}

