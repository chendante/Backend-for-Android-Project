package controller

import (
	"Backend-for-Android-Project/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func SignIn(c *gin.Context) {
	id := c.PostForm("id")
	password := c.PostForm("password")
	uid, _ := strconv.Atoi(id)
	name, userType, success := model.SignIn(uint(uid), password)
	var token string
	if success{
		token = model.CreatToken(uint(uid))
	}
	c.JSON(200, gin.H{
		"success":  success,
		"name": name,
		"type": userType,
		"token": token,
	})
}

func SignToken(c *gin.Context) {
	token := c.PostForm("token")
	name, userType, success := model.SignToken(token)
	if !success{
		token = ""
	}
	c.JSON(200, gin.H{
		"success":  success,
		"name": name,
		"type": userType,
		"token": token,
	})
}

func StudentInfo(c *gin.Context) {
	token := c.PostForm("token")
	id := model.Token2ID(token)
	if id == 0{
		token = ""
	}
	re, ok := model.StuSelectRegister(id)
	if !ok {
		re.Rid = 0
	}
	qu := model.SelectStudentQuestion(id)
	c.JSON(200, gin.H{
		"success":  ok,
		"register": re,
		"question": qu,
		"token": token,
	})
}