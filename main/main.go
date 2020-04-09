package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	_ "golang.org/x/crypto/ssh"
	"hackweek/main/db"
	"hackweek/main/router"
	"os"
)

func main() {
	InitConfig()
	db:=db.InitDB()
	defer db.Close()//关闭数据库
	r:=gin.Default()
    r=router.SetupRouter()
    port:=viper.GetString("server.port")
    if port !=""{
    	panic(r.Run(":"+port))
	}
	panic(r.Run())
}

func InitConfig()  {
	workDir,_:=os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir+"/config")
	err:=viper.ReadInConfig()
	if err!=nil{
		panic(err)
	}
}


