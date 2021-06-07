package maincontrollers
import (
	"net/http"
	// "fmt"
	"github.com/gin-gonic/gin"
	// "../mainmodels"
	. "../services/miscs"
	"github.com/robbert229/jwt"
	"encoding/json"
	"errors"
	"strings"
  )

type authController struct {}

var AuthController = authController{}

type testSuccess struct{
	Token string `json:"token"`
}

type authRequest struct{
	Username string 
	Password string  
	Test string
	
}

func (AC authController) Login(c *gin.Context ) {
	var req authRequest
	num, _ := c.GetRawData()
	json.Unmarshal(num, &req)
	
	if(req.Username != "username" && req.Password != "password"){
		// var er error = errors.New("wrong username & password")
		c.JSON(401, ErrorResponse{
			Iserror: 1,
			Msg: "wrong username & password",
		})
		return 
	}
	// userList,err := userModel.GetUsers()
	
	token, err := AC.generateToken(
		func(claims *jwt.Claims) {
			claims.Set("Role", "Admin")
			claims.Set("Username", req.Username)	
		})
	// token := "test"
	// var err error = nil
	

	if err != nil {
		var l = ErrorResponse{
			Iserror: 1,
			Msg: err.Error(),
		}

		c.JSON(500, l)
	}

	data := testSuccess{
		Token : token,
	}

	c.JSON(http.StatusOK, data)
}

type claimSetup func(*jwt.Claims)
func (AC authController) generateToken(f claimSetup) (string, error) {
	algorithm :=  jwt.HmacSha256(AC.GetSecretKey())
	claims := jwt.NewClaim()
	f(claims)
	token, err := algorithm.Encode(claims)
	return token, err
}

func (AC authController) GetSecretKey() string {
	return "HelloSecretKey"
}

func (AC authController) Decode(token string) (*jwt.Claims , error) {
	algorithm :=  jwt.HmacSha256(AC.GetSecretKey())
	dclaims, err := algorithm.Decode(token)
	return dclaims, err
}

func AuthenticationCheck(c *gin.Context , allowedUser []string) (*jwt.Claims, error){
	token := c.Request.Header["Authorization"]
	if(token == nil){
		return nil, errors.New("Token Not found")
	}
	//find bearer
	if(strings.Index(token[0], "Bearer ") == -1){
		return nil, errors.New("Bearer Token Not found")
	}

	parsedToken := (strings.Split(token[0] , "Bearer "))
	algorithm :=  jwt.HmacSha256(AuthController.GetSecretKey())
	v := algorithm.Validate(parsedToken[1])

	if(v != nil){
		return nil, errors.New("invalid token")
	}
	
	claims, err := algorithm.Decode(parsedToken[1])
	if(err == nil){
		return claims, nil
	}
	Var_dump("token halo", claims, err); 
	// Var_dump("token now", claims, err); 
	
	return nil, errors.New("Unauthenticated Token")
}