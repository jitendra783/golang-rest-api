package db

import "github.com/gin-gonic/gin"

type LogInterface interface {
	Registration(c *gin.Context, userId,password string)(string ,error)
	Login(c *gin.Context,userid, password string)(string,error)
	Logout(c *gin.Context, userid, password string)(string,error)
}
