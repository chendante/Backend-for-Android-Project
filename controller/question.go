package controller

import (
	"Backend-for-Android-Project/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CreateQuestion(c *gin.Context) {
	file, _ := c.FormFile("file")
	tmp := c.PostForm("lid")
	lid, _ := strconv.Atoi(tmp)
	token := c.PostForm("token")
	id := model.Token2ID(token)
	if id == 0{
		token = ""
	}
	uid := uint(lid)
	fileName, success := model.CreateQuestion(uid, file, c)
	c.JSON(200, gin.H{
		"success":  success,
		"fileName": fileName,
		"token": token,
	})
}

func UpdateQuestion(c *gin.Context) {
	token := c.PostForm("token")
	id := model.Token2ID(token)
	if id == 0{
		token = ""
	}
	qid := c.PostForm("qid")
	uid, _ := strconv.Atoi(qid)
	questionNum := c.PostForm("questionNum")
	tmp, _ := strconv.Atoi(questionNum)
	model.UpdateQuestion(uint(uid), uint(tmp))
	c.JSON(200, gin.H{
		"success":  true,
		"token": token,
	})
}

func GetQuestionDownloadUrl(c *gin.Context) {
	token := c.PostForm("token")
	id := model.Token2ID(token)
	if id == 0{
		token = ""
	}
	qid := c.PostForm("qid")
	uid, _ := strconv.Atoi(qid)
	fileUrl, ok := model.GetQuestionFileUrl(uint(uid))
	c.JSON(200, gin.H{
		"success":  ok,
		"token": token,
		"downloadUrl":fileUrl,
	})
}

func PostAnswer(c *gin.Context) {
	token := c.PostForm("token")
	id := model.Token2ID(token)
	if id == 0{
		token = ""
	}
	answer := c.PostForm("answer")
	qid := c.PostForm("qid")
	uid, _ := strconv.Atoi(qid)
	tmp := c.PostForm("questionNum")
	questionNum, _ := strconv.Atoi(tmp)
	ok := model.PostAnswer(id, uint(uid), uint(questionNum), answer)
	c.JSON(200, gin.H{
		"success":  ok,
		"token": token,
	})
}

func SelectAnswers(c *gin.Context) {
	token := c.PostForm("token")
	id := model.Token2ID(token)
	if id == 0{
		token = ""
	}
	qid := c.PostForm("qid")
	uid, _ := strconv.Atoi(qid)
	questionNum := c.PostForm("questionNum")
	tmp, _ := strconv.Atoi(questionNum)
	answers := model.SelectAnswers(uint(uid), uint(tmp))
	c.JSON(200, gin.H{
		"success":  true,
		"token": token,
		"answers":answers,
	})
}
