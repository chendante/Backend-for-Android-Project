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

type QuestionInfo struct {
	Qid	uint
	Lid uint
	QuestionNum uint
	FileName string
	FilePath string
	UniName  string
	LessonName	string
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

type AnswerInfo struct {
	Anid	uint
	Sid		uint
	Qid		uint
	QuestionNum	uint
	AnswerString string
	Name	string
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

func SelectAnswers(qid uint, questionNum uint) []AnswerInfo {
	var answerInfos []AnswerInfo
	//Db.Table("Questions").Select("Answers.*, user.name").Joins("inner Join StuLesson on StuLesson.lid = Questions.lid").Joins("inner Join user on StuLesson.sid = user.id").Joins("left Join Answers on Answers.sid = StuLesson.sid and Answers.qid = Questions.qid").Where("Answers.qid = ? AND Answers.question_num = ?", qid, questionNum)
	Db.Table("Answers").Select("Answers.*, user.name").Joins("inner Join user on Answers.sid = user.id").Where("Answers.qid = ? AND Answers.question_num = ?", qid, questionNum).Find(&answerInfos)
	return answerInfos
}

func SelectStudentQuestion(sid uint) QuestionInfo {
	var questionInfo QuestionInfo
	Db.Table("StuLesson").Select("Questions.* , Lesson.lesson_name").Joins("inner join Lesson on Lesson.lid = StuLesson.lid").Joins("inner Join Questions on Questions.lid = Lesson.lid").Where("StuLesson.sid = ? AND Questions.question_num != 0", sid).Last(&questionInfo)
	return questionInfo
}
