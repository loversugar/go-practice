package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	Name string `form:"name" json:"name"`
	Password string `form:"password" json:"password"`
}

func GetUserInfo(c *gin.Context) {
	var user User

	if c.ShouldBindQuery(&user) != nil {
		log.Print(user.Name)
	} else {
		log.Println(111)
		log.Println(user.Name)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func CreateUser(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"status": "error"})
		return
	} else {
		log.Println("success")
		log.Println(user.Name)
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
