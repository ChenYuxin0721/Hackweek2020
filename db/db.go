package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"hackweek/main/model"
)

var DB *gorm.DB


/*
func InitMySQL()(err error){
	dsn :="root:root@tcp(127.0.0.1:13306)/db1?charset=utf8&mb4&parseTime=True&loc=local"
	DB, err =gorm.Open("mysql",dsn)
	if err != nil{
		return
	}
	return DB.DB().Ping()
}
*/

func InitDB()*gorm.DB  {
	driverName:=viper.GetString("datasource.driverName")
	host:=viper.GetString("datasource.host")
	port:=viper.GetString("datasource.port")
	database:=viper.GetString("datasource.database")
	username:=viper.GetString("datasource.username")
	password:=viper.GetString("datasource.password")
	charset:=viper.GetString("datasource.charset")
	args:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err :=gorm.Open(driverName,args)
	if err !=nil{
		panic("failed to connect database,err:"+err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB=db
	return db
}

func GetDB()  *gorm.DB{
	return DB
}
