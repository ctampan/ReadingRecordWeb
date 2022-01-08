package auth

import (
	"fmt"

	"github.com/ctampan/ReadingRecordWeb/readingrecord-backend/jwtFunction"
	"github.com/gin-gonic/gin"
)

func ValidateDemoEndpoint(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(401, gin.H{})
	} else {
		username, err := jwtFunction.JwtValidateToken(token)
		if err != nil {
			fmt.Print(err.Error())
			c.JSON(401, gin.H{})
		} else {
			c.JSON(200, gin.H{
				"Username": username,
			})
		}
	}
}
