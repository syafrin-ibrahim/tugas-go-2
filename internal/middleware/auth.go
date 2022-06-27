package middleware

import (
	"context"
	"course/internal/user"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	bearer = "Bearer "
)

//type contextKey string

func Auth(user *user.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(403, gin.H{
				"message": "not authorize",
			})
			c.Abort()
		}
		fmt.Println(authHeader)
		fmt.Println("===========")
		fmt.Println(bearer)
		if !strings.HasPrefix(authHeader, bearer) {
			c.JSON(403, gin.H{
				"message": "not authorize 1",
			})
			c.Abort()
		}

		auths := strings.Split(authHeader, " ")
		fmt.Println(auths[1])
		data, err := user.DecriptJwt(auths[1])

		if err != nil {
			c.JSON(403, gin.H{
				"message": "not authorize 2",
			})
			c.Abort()
		}
		ctxUserId := context.WithValue(c.Request.Context(), "user_id", data["user_id"])
		c.Request = c.Request.WithContext(ctxUserId)
		c.Next()

	}
}
