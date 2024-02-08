package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	// Establish a database connection
	//db, err = gorm.Open("mysql", "root:Jitu@2790@tcp(127.0.0.1:3306)/testapi?charset=utf8&parseTime=True&loc=Local")

	user :=""
	password:=""
	host:= "localhost"
	port:= "3306"
	database_name := ""

	db, err := sql.Open("mysql", user+":"+password+"@tcp("+host+":" +port+")/"+database_name)
	if err != nil {
		return nil, err
	}

	// Ping the database to ensure the connection is successful
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database!")

	return db, nil
}
