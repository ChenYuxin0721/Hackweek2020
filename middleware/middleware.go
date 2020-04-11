package middleware

import (
	"github.com/gin-gonic/gin"
	"hackweek/main/db"
	"hackweek/main/model"
	"net/http"
	"strings"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取authorization header
		tokenString := c.GetHeader("Authorization")

		//validate token fomate
		if tokenString == " " || !strings.HasPrefix(tokenString, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    400,
				"message": "权限不足"})
			c.Abort()
			return
		}
		tokenString = tokenString[7:]

		token, claims, err := db.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    400,
				"message": "权限不足"})
			c.Abort()
			return
		}
		//验证通过后获取claim中的userId
		userId := claims.UserId
		DB := db.GetDB()
		var user model.User
		DB.First(&user, userId)
		//用户不存在
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    400,
				"message": "权限不足"})
			c.Abort()
			return
		}
		// 用户存在则将user信息写入上下文
		c.Set("user", user)

		c.Next()

	}
}

func CrossOrigin() gin.HandlerFunc {

	return func(c *gin.Context) {

		allowHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, x-CSRF-Token, Authorization"
		if origin := c.Request.Header.Get("Origin"); origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Vary", "Origin")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT, PATCH, GET, DELETE")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Headers", allowHeaders)
		}
		c.Request.Header.Del("Origin")
	}

}
