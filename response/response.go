package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(c *gin.Context, httpStatus int,code int,data gin.H,message string)  {
	c.JSON(httpStatus,gin.H{"code":code, "message":message, "data":data})
}

func Success(c *gin.Context,data gin.H,message string)  {
	Response(c,http.StatusOK,200,data,message)
}

func Fail(c *gin.Context,data gin.H,message string)  {
	Response(c,http.StatusOK,400,data,message)
}

