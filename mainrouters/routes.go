//routes.go
package mainrouters

import (
  "github.com/gin-gonic/gin"
  mc "../maincontrollers"
  md "../middleware"
)


func SetupRoute(router *gin.Engine){
    router.GET("/", mc.ShowIndexPage)
    router.GET("/users", mc.GetUsers)
    router.POST("/auth", mc.AuthController.Login)
    router.GET("/user/:id",md.ForLoggedInUser(), mc.UserController.GetUserByID)
    router.POST("/user/", mc.UserController.GetUserByID)
    router.PUT("/user/:id", mc.UserController.GetUserByID)
    router.DELETE("/user/:id", mc.UserController.GetUserByID)
    router.DELETE("/user/:id/confirm", mc.UserController.GetUserByID)
}

