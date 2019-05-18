package main

import "Backend-for-Android-Project/router"

func main() {
	r := router.InitRouter()
	_ = r.Run("0.0.0.0:8000")
}