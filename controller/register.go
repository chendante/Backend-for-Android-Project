package controller

import (
	"Backend-for-Android-Project/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func NewRegister(c *gin.Context) {
	var rid int
	token, _ := c.GetPostForm("token")
	lid, _ := c.GetPostForm("lid")
	id := model.Token2ID(token)
	if id == 0{
		token = ""
	} else {
		uid, _ := strconv.Atoi(lid)
		rid = model.CreateRegister(uint(uid))
	}
	c.JSON(200, gin.H{
		"rid": rid,
		"token": token,
	})
}

func DeleteRegister(c *gin.Context) {
	token, _ := c.GetPostForm("token")
	rid, _ := c.GetPostForm("rid")
	id := model.Token2ID(token)
	if id == 0{
		token = ""
	} else {
		uid, _ := strconv.Atoi(rid)
		model.DeleteRegister(uint(uid))
	}
	c.JSON(200, gin.H{
		"token": token,
	})
}

func GetAttendance(c *gin.Context) {
	token, _ := c.GetPostForm("token")
	rid, _ := c.GetPostForm("rid")
	id := model.Token2ID(token)
	var attendances []model.AttendanceBook
	if id == 0{
		token = ""
	} else {
		uid, _ := strconv.Atoi(rid)
		attendances = model.SelectAttendance(uint(uid))
	}
	c.JSON(200, gin.H{
		"token": token,
		"attendances": attendances,
	})
}

func StuAttend(c *gin.Context) {
	token, _ := c.GetPostForm("token")
	rid, _ := c.GetPostForm("rid")
	id := model.Token2ID(token)
	var success = false
	if id == 0{
		token = ""
	} else {
		uid, _ := strconv.Atoi(rid)
		success = model.PostAttendance(id, uint(uid))
	}
	c.JSON(200, gin.H{
		"token": token,
		"success": success,
	})
}
