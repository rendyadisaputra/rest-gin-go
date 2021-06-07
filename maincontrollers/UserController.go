package maincontrollers
import (
	"net/http"
	// "fmt"
	"github.com/gin-gonic/gin"
	"../mainmodels"
	. "../services/miscs"
	"plugin"
  )
  
var router *gin.Engine

type userController struct {}

var UserController = userController{}

func LoadModule(pluginName string ) (*plugin.Plugin, error) {
	pluginModule, err := plugin.Open("plugins/"+pluginName+"/"+pluginName+".so")
	if( err != nil ){
		return nil, err	
	}
	return pluginModule, nil
}

func (UC userController) GetUserByID(c *gin.Context ) {
	
	/** testing load module */
	pluginM, _ := LoadModule("firstplugin")
	printHelloWorldSymbol, errP := pluginM.Lookup("PrintHelloWorld")
	printHelloWorldFunc := printHelloWorldSymbol.(func())
	printHelloWorldFunc()
	Var_dump("TYPE", errP)

	// printHelloWorld()

	userModel := mainmodels.UserModel()
	userList,err := userModel.GetUsers()
	Var_dump("here test")

	if (err != nil){

		var l = ErrorResponse{
			Iserror: 1,
			Msg: err.Error(),
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
			Msg: err.Error(),
		}
		Var_dump("it is error")

		
		c.JSON(500, l)
	} else{
		Var_dump(userList)
		mainmodels.MysqlClose()
		c.JSON(http.StatusOK, userList)
	}
}