// main.go

package main

import (
  "github.com/gin-gonic/gin"
  "./Routers"
  "io/ioutil"
)
import (
  "database/sql"
_ "github.com/go-sql-driver/mysql")

var router *gin.Engine

func main() {

  // Set the router as the default one provided by Gin
  gin.SetMode(gin.ReleaseMode)
  gin.DefaultWriter = ioutil.Discard

  router = gin.Default()
  
  router.Static("/assets", "./assets")
  // Process the templates at the start so that they don't have to be loaded
  // from the disk again. This makes serving HTML pages very fast.
  router.LoadHTMLGlob("View/templates/*")

  mainrouter.SetupRoute(router);
  // Define the route for the index page and display the index.html template
  // To start with, we'll use an inline route handler. Later on, we'll create
  // standalone functions that will be used as route handlers.
  
  mysql_test()
  // Start serving the application
  router.Run(":8000")

}


func mysql_test(){

    // Open up our database connection.
    // I've set up a database on my local machine using phpmyadmin.
    // The database is called testDb
    db, err := sql.Open("mysql", "root:mypass@tcp(127.0.0.1:3306)/n14580dl_myskinwebsite-local")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    // defer the close till after the main function has finished
    // executing
    defer db.Close()
}