package api

import (
	"example/pkg/config"
	"example/pkg/db"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func Server() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	flag.Parse()
	config.LoadConfig(*environment)
	log.Println("server start")
	config := config.GetConfig()
	r := Router()
	r.Run(config.GetString("server.port"))
}

func custom_logger(c *gin.Context) {

}
