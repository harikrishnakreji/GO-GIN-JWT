package main

import (
	"github.com/gin-gonic/gin"
	"github.com/harikrishnakreji/GO-JWT/controllers"
	"github.com/harikrishnakreji/GO-JWT/initializers"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDb()
	initializers.SyncDataBase()
}
func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.Run()
}
