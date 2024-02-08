package v1

import (
	"log"
    "net/http"
	"github.com/gin-gonic/gin"
	"example/pkg/db"
)
type LogController struct{
	
}
var LoginRes =new(db.LoginModel)
type User struct{
	UserID string `json:"userId"`
	Password string `json:"password"`
} 

func( l LogController) Login(c *gin.Context){
	log.Println("registration ")
	var user User
	// Bind the request data to the User struct
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    response,err :=LoginRes.Login(c, user.UserID,user.Password)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK,gin.H{"massage": "success", "data": response})
}

func( l LogController) Registration(c *gin.Context){
	log.Println("login ")
	var user User
	// Bind the request data to the User struct
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    response,err :=LoginRes.Registration(c, user.UserID,user.Password)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK,gin.H{"massage": "success", "data": response})
}

func( l LogController) Logout(c *gin.Context){
	log.Println("logout ")
	var user User
	// Bind the request data to the User struct
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    response,err :=LoginRes.Logout(c, user.UserID,user.Password)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK,gin.H{"massage": "success", "data": response})
}