package main

import (
	"example/Config"
<<<<<<< HEAD
	"example/api"
=======
	"example/Models"
	"example/Routers"
>>>>>>> 7d875aa070a302c8129e6a1517c4e59627317d96
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

func main() {

	Config.DB, err = gorm.Open("mysql", "root:Jitu@2790@tcp(127.0.0.1:3306)/testapi?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("statuse: dataabsae not connect ")
	}
	defer Config.DB.Close()
<<<<<<< HEAD
	api.Server()
=======
	Config.DB.AutoMigrate(&Models.User{})

	r := Routers.SetupRouter()
	// running
	r.Run()
>>>>>>> 7d875aa070a302c8129e6a1517c4e59627317d96
}
