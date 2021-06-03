package controller
import (
	"net/http"
	// "fmt"
	"github.com/gin-gonic/gin"
	"../Models/structs"
  )
  
  
func ShowIndexPage(c *gin.Context ){
	// fmt.Println(fmt.Sprintf("print broh  %#v", c))
	  // Call the HTML method of the Context to render a template
	//   c.HTML(http.StatusOK, "index.html", gin.H{
	// 	  "test" : "title",
	// 	  "title" : "homepage",
	//   })	

	c.JSON(http.StatusOK, structs.ArticleList)
  
  }