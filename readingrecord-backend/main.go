package main

import (
	"github.com/ctampan/ReadingRecordWeb/readingrecord-backend/endpoint/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/login", auth.LoginEndpoint)
	r.POST("/validateDemo", auth.ValidateDemoEndpoint)
	r.Run(":8000")
}
