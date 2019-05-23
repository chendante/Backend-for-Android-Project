package model

import (
	. "Backend-for-Android-Project/model/base"
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"strconv"
	"time"
)

type Lesson struct {
	Lid        	uint    	`gorm:"primary_key"`
	LessonName 	string  	`gorm:"size:255"`
	LessonTime 	string  	`gorm:"size:255"`
	Tid			uint
	LessonAddress	string
	Name		string
}

func (Lesson) TableName() string {
	return "Lesson"
}

type StuLesson struct {
	Sid		uint
	Lid		uint
	SLID	uint	`gorm:"primary_key"`
}

func (StuLesson) TableName() string {
	return "StuLesson"
}

type Attachment struct {
	Aid      uint
	FileName string
	FilePath string
	UniName  string
	Lid      uint
}

func (Attachment) TableName() string {
	return "Attachment"
}

func GetStudentLessons(SID uint) ([]Lesson, bool) {
	var lessons []Lesson
	Db.Table("StuLesson").Select("Lesson.*, user.name").Joins("inner join Lesson on Lesson.lid = StuLesson.lid").Joins("inner join user on user.id = Lesson.tid").Where("StuLesson.sid = ?", SID).Find(&lessons)
	if len(lessons) == 0 {
		return nil, false
	}else {
		return lessons, true
	}
}

func GetTeacherLessons(TID uint) ([]Lesson, bool) {
	var lessons []Lesson
	Db.Table("Lesson").Select("Lesson.*").Where(Lesson{Tid:TID}).Find(&lessons)
	if len(lessons) == 0{
		return nil, false
	}else {
		return lessons, true
	}
}

func UploadAttachment(file *multipart.FileHeader, c *gin.Context, tid uint, lid uint) (string, bool) {
	if tid == 0{
		return "", false
	}
	uniName := file.Filename
	fileName := file.Filename
	filePath := UploadUri + uniName
	err := c.SaveUploadedFile(file, filePath)
	attachment := Attachment{FileName: fileName, FilePath: filePath, Lid: lid, UniName: uniName}
	Db.Create(&attachment)
	var success = true
	fmt.Println(attachment.Aid)
	if !Db.NewRecord(&attachment){
		success = false
	}
	if err != nil{
		success = false
		fmt.Println(err.Error())
	}
	return fileName, success
}

func GetAttachmentUniName(lid uint) (string, bool) {
	var attachment Attachment
	rowNum := Db.Where(&Attachment{Lid:lid}).First(&attachment).RowsAffected
	if rowNum == 0{
		return "", false
	} else {
		return attachment.UniName, true
	}
}

func DownloadUrl(uniName string) string {
	return BaseUrl+uniName
}

func getUniName() string {
	cruTime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(cruTime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token
}
