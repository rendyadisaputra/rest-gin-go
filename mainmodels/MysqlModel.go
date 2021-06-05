package mainmodels

import (
  "database/sql"
_ "github.com/go-sql-driver/mysql"
  "github.com/joho/godotenv"
  "os"
  "../services/miscs"
)

var db  *sql.DB

func MysqlModel() *sql.DB{

    // Open up our database connection.
    // The database is called testDb
    err := godotenv.Load() // it will load .env in the same directory by default;
    mysql_user := os.Getenv("MYSQL_USER");
    mysql_pass := os.Getenv("MYSQL_PASSWORD");
    mysql_host := os.Getenv("MYSQL_HOST");
    mysql_dbname := os.Getenv("MYSQL_DBNAME");
    

    if err != nil {
      panic(err.Error())
    }


    db, err := sql.Open("mysql", mysql_user+":"+mysql_pass+"@tcp("+mysql_host+")/"+mysql_dbname)

    miscs.Var_dump(db)
    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    // defer the close till after the main function has finished
    // executing
   
    

  return db
}

func MysqlClose(){
	db.Close()
}