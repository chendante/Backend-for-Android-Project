package router

import "github.com/gin-gonic/gin"
import . "Backend-for-Android-Project/controller"

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/sign-in", SignIn)
	router.POST("/sign-token", SignToken)
	router.POST("/student/get-lesson", GetStuLessons)
	router.POST("/teacher/new-register", NewRegister)
	router.POST("/teacher/delete-register", DeleteRegister)
	router.POST("/teacher/get-attendance", GetAttendance)
	router.POST("/student/attend", StuAttend)
	return router
}
