package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/harikrishnakreji/GO-JWT/initializers"
	"github.com/harikrishnakreji/GO-JWT/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	//Get the Email/Password of Req Body
	var body struct {
		Email    string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	//Hash the Password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
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
func Login(c *gin.Context) {
	//Get the email and pass off req body
	var body struct {
		Email    string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	//Look up requested user
	var user models.User
	initializers.DB.First(&user, "email", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email or password",
		})
		return
	}

	//comapre sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email or password",
		})
		return
	}
	//Generate a jwt token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(os.Getenv("SECRET"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "FAIL TO CREATE TOKEN",
		})
		return
	}
	//send it back
	c.JSON(http.StatusBadRequest, gin.H{
		"token": tokenString,
	})

}
