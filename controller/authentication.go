package controller

import (
	//"diaryApi/helper"
	"diaryApi/model"
	//"errors"
	//"fmt"
	"net/http"

	//"os/user"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input model.AuthenticationInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest , gin.H{"error": err.Error() })
		return
	}

	user := model.User{
		Username: input.Username,
		Password: input.Password,
	}

	savedUser , err := user.Save()
	
	if err != nil {
		c.JSON(http.StatusBadRequest , gin.H{"error" : err.Error()})
		return
	}
	c.JSON(http.StatusCreated , gin.H{"user" : savedUser})	
}

func Login(c *gin.Context) {
	var input model.AuthenticationInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest , gin.H{"error" : err.Error()})
		return
	}

	user , err := model.FindUserByUsername(input.Username)
	
	if err != nil {
		c.JSON(http.StatusBadRequest , gin.H{"error" : err.Error()})
		return
	}
	user.ValidatePassword(input.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest , gin.H{"1error" : err.Error()})
		return
	}

	// jwt , err := helper.GenerateJWT(user)

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest , gin.H{"2error" : err.Error()})
	// 	return
	// }

	// c.JSON(http.StatusOK , gin.H{"jwt" : jwt})
	c.JSON(http.StatusOK , user)

}

func GetUsers (c *gin.Context) {
	users , err := model.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusNotFound , gin.H{"Error" : err.Error()})
	}

	c.JSON(http.StatusOK , users)

}