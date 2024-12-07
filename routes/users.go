package routes

import (
	"Rest-API/models"
	"Rest-API/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func signUp(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(400, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(400, gin.H{"message": "Could not save user"})
	}

	context.JSON(200, gin.H{"message": "User saved successfully."})

}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(400, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.ValidateLoginCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		fmt.Print("Validate: ", err.Error())
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		fmt.Print("Due to token: ", err.Error())
		return
	}

	context.JSON(200, gin.H{"message": "Login successful.", "token": token})
}
