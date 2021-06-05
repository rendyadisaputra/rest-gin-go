package mainmodels

import (
	"../services/miscs"
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"

)

type User struct {
	ID   int    `json:"id"`
	Email string `json:"user_email"`
}

type userMethod struct{
	TableName string
}

func (uM userMethod) GetTableName() string{
	return uM.TableName;
}

func UserModel() userMethod{
	db =  MysqlModel();
	return  userMethod{
		TableName: "wp_users",
	};
}

func (uM userMethod) GetUsers() ([]User, error){
	var userList []User

	results, err := db.Query("SELECT ID, user_email FROM "+uM.TableName+" LIMIT 10")

    if err != nil {
	  miscs.Var_dump(err)
      return userList, err // proper error handling instead of panic in your app
    }

	
	for results.Next(){
		var user User
		err := results.Scan(&user.ID, &user.Email)
		if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
             
		userList = append(userList, user)
		miscs.Var_dump(user.ID)
	}
	 return userList, err
}

