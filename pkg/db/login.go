package db

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (l LoginModel) Registration(c *gin.Context, userid, password string) (string, error) {
	log.Println("login model")

	var err error
	return "success", err
}

func (l LoginModel) Login(c *gin.Context, userid, password string) (string, error) {
	log.Println("login model")

	var err error
	return "success", err
}

func (l LoginModel) Logout(c *gin.Context, userid, password string) (string, error) {
	log.Println("login model")

	var err error
	return "success", err
}
