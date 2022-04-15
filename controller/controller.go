package controller

import (
	"JeffMusic/models"
	"JeffMusic/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
 url     --> controller  --> logic   -->    model
请求来了  -->  控制器      --> 业务逻辑  --> 模型层的增删改查
*/

func RegisterNewAccount(c *gin.Context) {
	// 1. 从请求中把数据拿出来
	var user models.User
	c.BindJSON(&user)
	// 2. 存入数据库
	err := models.CreateNewAccount(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func Login(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	user, err := models.GetUserInfo(u.UnionID)
	if err != nil || user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	token, err := utils.CreateToken(user.UnionID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}

func ValidateToken(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusForbidden, "The authorization token is abnormal.")
	}
	token, err := utils.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "invalid token")
		return
	}
	c.JSON(http.StatusOK, token)
}
