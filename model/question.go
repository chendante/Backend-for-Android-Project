package model

import (
	. "Backend-for-Android-Project/model/base"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type Question struct {
	Qid	uint	`gorm:"primary_key"`
	Lid uint
	QuestionNum uint
	FileName string
	FilePath string
	UniName  string
}

func (Question) TableName() string {
	return "Questions"
}

type Answer struct {
	Anid	uint
	Sid		uint
	Qid		uint
	QuestionNum	uint
	AnswerString string
}

func (Answer) TableName() string {
	return "Answers"
}

func CreateQuestion(lid uint, file *multipart.FileHeader, c *gin.Context) (string, bool) {
	var success = true
	uniName := getUniName()
	fileName := file.Filename
	filePath := UploadUri + uniName
	err := c.SaveUploadedFile(file, filePath)
	if err != nil{
		success = false
		fmt.Println(err.Error())
	}
	question := Question{Lid: lid, QuestionNum:0, FileName:fileName, FilePath:filePath, UniName:uniName}
	Db.Create(&question)

	return fileName, success
}

func UpdateQuestion(Qid uint, questionNum uint) {
	var question Question
	Db.First(&question, Qid)
	question.QuestionNum = questionNum
	Db.Save(&question)
}

func GetQuestionFileUrl(Qid uint) (string, bool) {
	var question Question
	rowNum := Db.Where(&Question{Qid:Qid}).First(&question).RowsAffected
	if rowNum == 0{
		return "", false
	} else {
		return DownloadUrl(question.UniName), true
	}
}

func PostAnswer(sid uint, qid uint, questionNum uint, answerString string) bool {
	answer := Answer{Sid:sid, Qid:qid, QuestionNum:questionNum, AnswerString:answerString}
	Db.Create(&answer)
	return !Db.NewRecord(answer)
}

//func SelectAnswers(qid uint, questionNum uint) []Answer {
//
//}

//func SelectStudentQuestion(sid uint) {
//
//}
