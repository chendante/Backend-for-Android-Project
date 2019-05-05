package main

import "Backend-for-Android-Project/router"

func main() {
	r := router.InitRouter()
	_ = r.Run("localhost:8000")
}