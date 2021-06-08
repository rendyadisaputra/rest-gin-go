package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	mc "../maincontrollers"
	. "../services/miscs"
	"errors"
	"github.com/robbert229/jwt"
	"strings"

)

func ForLoggedInUser() gin.HandlerFunc{

	return func(c *gin.Context){
		_, err := (AuthenticationCheck(c))
		if(err != nil){
			c.AbortWithStatusJSON(http.StatusUnauthorized, mc.ErrorResponse{
				Iserror: 1,
				Msg: "Unauthenticated",
			})
		}
	}
}

func AuthenticationCheck(c *gin.Context) (*jwt.Claims, error){
	token := c.Request.Header["Authorization"]
	if(token == nil){
		return nil, errors.New("Token Not found")
	}
	//find bearer
	if(strings.Index(token[0], "Bearer ") == -1){
		return nil, errors.New("Bearer Token Not found")
	}

	parsedToken := (strings.Split(token[0] , "Bearer "))
	algorithm :=  jwt.HmacSha256(mc.AuthController.GetSecretKey())
	v := algorithm.Validate(parsedToken[1])

	if(v != nil){
		return nil, errors.New("invalid token")
	}
	
	claims, err := algorithm.Decode(parsedToken[1])
	
	if(err == nil){
		role, _ := claims.Get("Role");
		c.Set("Role", role)

		username, _ := claims.Get("Username");
		c.Set("Username", username)
		return claims, nil
	}
	
	Var_dump("token now", claims, err); 
	
	return nil, errors.New("Unauthenticated Token")
}