package routes

import (
	"net/http"

	"github.com/JoaoFerrareis02/REST-API-GO/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {

	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data."})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User saved successfully!"})

}

func login(context *gin.Context) {

	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data."})
		return
	}

}
