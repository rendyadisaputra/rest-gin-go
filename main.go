package main
// main.go

import (
  "github.com/gin-gonic/gin"
  "./mainrouters"
  // "io/ioutil"
)

func main() {

  // Set the router as the default one provided by Gin
  /** Part of Release mode don't delete it */
  // gin.SetMode(gin.ReleaseMode)
  // gin.DefaultWriter = ioutil.Discard
  /** END OF RELEASE MODE */

 var router = gin.Default()
  
  router.Static("/assets", "./publics/assets")
  // Process the templates at the start so that they don't have to be loaded
  // from the disk again. This makes serving HTML pages very fast.
  router.LoadHTMLGlob("view/templates/*")

  mainrouters.SetupRoute(router);
  // Define the route for the index page and display the index.html template
  // To start with, we'll use an inline route handler. Later on, we'll create
  // standalone functions that will be used as route handlers.
  
  // Start serving the application
  router.Run(":8000")

}


