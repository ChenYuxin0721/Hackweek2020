package router

import (
	"github.com/gin-gonic/gin"
	"hackweek/main/controller"
	"hackweek/main/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CrossOrigin())
	//api
	v1Group := r.Group("v1")
	{
		//注册
		v1Group.POST("/api/register", controller.Register)
		//登录
		v1Group.POST("/api/login", controller.Login)
		//存入信息
		v1Group.GET("/api/info", middleware.Middleware(), controller.Info)
		//发表故事
		v1Group.POST("/api/create", controller.CreateAStory)
		//查看我的所有的故事
		v1Group.GET("/api/readme", controller.ReadAllMyStory)
		//修改我的某个故事
		v1Group.PUT("/api/readme/:id", controller.UpdateAStory)
		//删除我的某个故事
		v1Group.DELETE("/api/readme", controller.DeleteAStory)
	}

	v2Group := r.Group("/v2") //数据库添加用户组
	{
		//删除用户
		v2Group.DELETE("/:id", controller.DeleteUser)

	}
	return r
}
