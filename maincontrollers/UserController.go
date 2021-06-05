package maincontrollers
import (
	"net/http"
	// "fmt"
	"github.com/gin-gonic/gin"
	"../mainmodels"
	. "../services/miscs"
  )
  
var router *gin.Engine

type userController struct {}

var UserController = userController{}

func (UC userController) GetUserByID(c *gin.Context ) {
	userModel := mainmodels.UserModel()
	userList,err := userModel.GetUsers()
	Var_dump("here test")

	if (err != nil){

		var l = ErrorResponse{
			Iserror: 1,
			Msg: err,
		}
		Var_dump("it is error")

		
		c.JSON(500, l)
	} else{
		Var_dump(userList)
		mainmodels.MysqlClose()
		c.JSON(http.StatusOK, userList)
	}
}

func GetUsers(c *gin.Context ) {
	/** Just for test */	
	//Lets Load the Model First
	userModel := mainmodels.UserModel()
	userList,err := userModel.GetUsers()
	Var_dump("here test")

	if (err != nil){

		var l = ErrorResponse{
			Iserror: 1,
			Msg: err,
		}
		Var_dump("it is error")

		
		c.JSON(500, l)
	} else{
		Var_dump(userList)
		mainmodels.MysqlClose()
		c.JSON(http.StatusOK, userList)
	}
}