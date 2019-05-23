package router

import (
	"github.com/gin-gonic/gin"
)
import . "Backend-for-Android-Project/controller"

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/sign-in", SignIn)
	router.POST("/sign-token", SignToken)
	router.POST("/student/get-lesson", GetStuLessons)
	router.POST("/teacher/get-lesson", GetTchLessons)
	router.POST("/lesson/upload-attachment", UploadAttachment)
	router.POST("/lesson/download-url", GetDownloadUrl)
	router.POST("/teacher/new-register", NewRegister)
	router.POST("/teacher/delete-register", DeleteRegister)
	router.POST("/teacher/get-attendance", GetAttendance)
	router.POST("/student/attend", StuAttend)
	router.POST("/teacher/create-question", CreateQuestion)
	router.POST("/teacher/update-question", UpdateQuestion)
	router.POST("/question/get-file-url", GetQuestionDownloadUrl)
	router.POST("/student/post-answer", PostAnswer)

	router.POST("/student/info", StudentInfo)
	router.POST("/teacher/answers", SelectAnswers)
	return router
}
