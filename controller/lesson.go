package controller

import (
	"Backend-for-Android-Project/model"
	"github.com/gin-gonic/gin"
	"net/http"
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

func GetDownloadUrl(c *gin.Context) {
	token := c.PostForm("token")
	id := model.Token2ID(token)
	if id == 0{
		token = ""
		c.JSON(200, gin.H{
			"token": token,
		})
	} else {
		tmp := c.PostForm("lid")
		lid, _ := strconv.Atoi(tmp)
		uid := uint(lid)
		uniName, success := model.GetAttachmentUniName(uid)
		var downloadUrl string
		if success{
			downloadUrl = model.DownloadUrl(uniName)
		}
		c.JSON(http.StatusOK, gin.H{
			"token": token,
			"success": success,
			"downloadUrl": downloadUrl,
		})
	}
}
