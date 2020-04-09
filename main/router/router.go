package router

import (
	"github.com/gin-gonic/gin"
	"hackweek/main/controller"
	"hackweek/main/middleware"
)

func SetupRouter()*gin.Engine{
	r := gin.Default()
	r.Static("/static","static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)
	//api
	v1Group := r.Group("v1")
	{
		//注册
		v1Group.POST("/register",controller.Register)
		//登录
		v1Group.POST("/login",controller.Login)
		//存入信息
		v1Group.GET("/info",middleware.Middleware(),controller.Info)
		//发表故事
		v1Group.POST("/post", controller.CreateAStory)
		//查看我的所有的故事
		v1Group.GET("/readme", controller.ReadAllMyStory)
		//查看我的某个故事
		v1Group.GET("/readme/:id", func(c *gin.Context) {

		})
		//修改我的某个故事
		v1Group.PUT("/readme/:id", controller.UpdateAStory)
		//删除我的某个故事
		v1Group.DELETE("/readme/:id", controller.DeleteAStory)
	}
	return r
}