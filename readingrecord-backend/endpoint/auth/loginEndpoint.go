package auth

import (
	"github.com/ctampan/ReadingRecordWeb/readingrecord-backend/jwtFunction"
	"github.com/gin-gonic/gin"
)

func LoginEndpoint(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")

	signedToken, err := jwtFunction.JwtGenerateToken(username, password)

	if err != nil {
		c.JSON(401, gin.H{})
	} else {
		c.JSON(200, gin.H{
			"access_token": signedToken,
		})
	}
}
