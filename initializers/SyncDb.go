package initializers

import "github.com/harikrishnakreji/GO-JWT/models"

func SyncDataBase() {
	DB.AutoMigrate(&models.User{})
}
