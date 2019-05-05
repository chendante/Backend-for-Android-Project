package controller

import (
	"Backend-for-Android-Project/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetStuLessons(c *gin.Context) {
	id := c.PostForm("id")
	uid, _ := strconv.Atoi(id)
	lessons, success := model.GetStudentLessons(uint(uid))
	c.JSON(200, gin.H{
		"success":  success,
		"lessons": lessons,
	})
}
