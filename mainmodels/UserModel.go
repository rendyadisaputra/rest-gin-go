package mainmodels

import (
	. "../services/miscs"
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"

)

/** STRUCTS */
type User struct {
	ID   int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type userMethod struct{
	TableName string
}


/** Public Functions */

func UserModel() userMethod{
	db =  MysqlModel();
	return  userMethod{
		TableName: "Users",
	};
}


/** User Methods */

func (uM userMethod) GetTableName() string{
	return uM.TableName;
}


func (uM userMethod) GetUsers() ([]User, error){
	var userList []User

	results, err := db.Query("SELECT ID, user_email FROM "+uM.TableName+" LIMIT 10")

    if err != nil {
	  Var_dump(err)
      return userList, err // proper error handling instead of panic in your app
    }

	
	for results.Next(){
		var user User
		err := results.Scan(&user.ID, &user.Username)
		if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
             
		userList = append(userList, user)
		Var_dump(user.ID)
	}
	 return userList, err
}

func (uM userMethod) FindByUsername( username string) ([]User, error){
	var userList []User

	results, err := db.Query(
		`
		SELECT ID, username, password FROM `+uM.TableName+` 
		WHERE username="`+username+`"
		LIMIT 1
		`,
	)

    if err != nil {
      return userList, err // proper error handling instead of panic in your app
    }

	for results.Next(){
		var user User
		err := results.Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
             
		userList = append(userList, user)
		Var_dump(user.ID)
	}
	 return userList, err
}

func (uM userMethod) Create( user User) (int, error){

	query := `
		INSERT INTO `+uM.TableName+` (username, password)
		VALUES ("`+user.Username+`","`+user.Password+`")
	`
	_, err := db.Query(query)

	
    if err != nil {
		Var_dump("ERror ", err.Error(), query)
      return 0, err // proper error handling instead of panic in your app
    }

	return 1, err
}
