package v1

import "github.com/gin-gonic/gin"

type LogInterface interface {
	Registration(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
}