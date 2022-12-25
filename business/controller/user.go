package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"redrock-test/business/services"
)

func Enroll(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"msg": "wrong format",
		})
		return
	}
	if err := services.Enroll(username, password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "create a user failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}

func Login(c *gin.Context) {
	//username := c.PostForm("username")
	//password := c.PostForm("password")

}
