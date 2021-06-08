package maincontrollers
import (
	"net/http"
	// "fmt"
	"github.com/gin-gonic/gin"
	model "../mainmodels"
	. "../services/miscs"
	"plugin"
	"encoding/json"
	bc "golang.org/x/crypto/bcrypt"
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


	
func HashPassword(password string) (string, error) {
    bytes, err := bc.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func (UC userController) GetUserByID(c *gin.Context ) {
	
	/** testing load module */
	pluginM, _ := LoadModule("firstplugin")
	printHelloWorldSymbol, _ := pluginM.Lookup("PrintHelloWorld")
	printHelloWorldFunc := printHelloWorldSymbol.(func())
	printHelloWorldFunc()
	Var_dump(c.Get("Role"))
	Var_dump("VR")

	// printHelloWorld()

	userModel := model.UserModel()
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
		model.MysqlClose()
		c.JSON(http.StatusOK, userList)
	}
}

type registerJson struct{
	Username string
	Password string
}


func (UC userController) Register(c *gin.Context ) {
	
	var dt registerJson
	num, _ := c.GetRawData()

	json.Unmarshal(num, &dt)

	if(dt.Username == "" || dt.Password == ""){
		r := ErrorResponse{
			Iserror: 1,
			Msg: "Need username and password",
		}
		 c.JSON(403,r )
		 return
	}

	userModel := model.UserModel()
	list , err := userModel.FindByUsername(dt.Username)
	
	if(err != nil){
		Var_dump("list ", err.Error())
	}

	if(len(list) != 0){
		r := ErrorResponse{
			Iserror: 1,
			Msg: "username exists",
		}
		c.JSON(403, r )
		return 
	}

	pwd , err := HashPassword( dt.Password)
	userStruct := model.User{
		Username: dt.Username,
		Password: pwd,
	}
	userModel.Create(userStruct)

	type responses struct{
		Success int `json:"success"`
	}

	succcessResponse := responses{
		Success: 1,
	}
	c.JSON(200, succcessResponse)
	return 
}

func GetUsers(c *gin.Context ) {
	/** Just for test */	
	//Lets Load the Model First
	userModel := model.UserModel()
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
		model.MysqlClose()
		c.JSON(http.StatusOK, userList)
	}
}