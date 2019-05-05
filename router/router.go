package router

import "github.com/gin-gonic/gin"
import . "Backend-for-Android-Project/controller"

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/sign-in", SignIn)
	router.POST("/student/get-lesson", GetStuLessons)
	return router
}
