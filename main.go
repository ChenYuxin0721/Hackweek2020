package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	_ "golang.org/x/crypto/ssh"
	"hackweek/main/db"
	"hackweek/main/router"
	"net/http"
	"os"
)

func main() {
	InitConfig()
	db := db.InitDB()
	defer db.Close() //关闭数据库
	r := gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找

	r.Static("/static", "static")

	// 告诉gin框架去哪里找模板文件

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", nil)

	})
	r = router.SetupRouter()
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
