package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harikrishnakreji/GO-JWT/initializers"
	"github.com/harikrishnakreji/GO-JWT/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	//Get the Email/Password of Req Body
	var body struct {
		Email    string
		password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	//Hash the Password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	//Vreate the User
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
	}

	//Respond
	c.JSON(http.StatusBadRequest, gin.H{})
}
