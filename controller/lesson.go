package controller

import (
	"Backend-for-Android-Project/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetStuLessons(c *gin.Context) {
	token := c.PostForm("token")
	id := model.Token2ID(token)
	if id == 0{
		token = ""
	}
	lessons, success := model.GetStudentLessons(id)
	c.JSON(200, gin.H{
		"success":  success,
		"lessons": lessons,
		"token": token,
	})
}

func GetTchLessons(c *gin.Context) {
	token := c.PostForm("token")
	id := model.Token2ID(token)
	if id == 0{
		token = ""
	}
	lessons, success := model.GetTeacherLessons(id)
	c.JSON(200, gin.H{
		"success":  success,
		"lessons": lessons,
		"token": token,
	})
}

func UploadAttachment(c *gin.Context) {
	file, _ := c.FormFile("file")
	tmp := c.PostForm("lid")
	lid, _ := strconv.Atoi(tmp)
	token := c.PostForm("token")
	id := model.Token2ID(token)
	if id == 0{
		token = ""
	}
	uid := uint(lid)
	fileName, success := model.UploadAttachment(file, c,id ,uid)
	c.JSON(200, gin.H{
		"success":  success,
		"fileName": fileName,
		"token": token,
	})
}
