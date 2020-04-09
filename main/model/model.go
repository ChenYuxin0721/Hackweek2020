package model

import (
	"github.com/jinzhu/gorm"
	"hackweek/main/db"
)

type User struct {
	gorm.Model
	name string `jason:"name"`
	password string `jason:"password"`
	following int64 `json:"'following's id"`
	follower int64 `json:"'follower'id"`
	like int64 `json:"like"`
	story_id uint64 `gorm:"association_foreignkey:Story;foreignkey:id"`
	currency int64
}


type Story struct{
	gorm.Model
	story string `json:"story"`
}

//增删查改
func CreateAStory() (story []*Story,err error) {
	err = db.DB.Create(&story).Error
	return
}

func ReadAllMyStory()(allstory []*Story,err error){
	if err := db.DB.Find(&allstory).Error;err!=nil{
		return  nil, err
	}
	return
}

func GETAStory(id string)(story *Story,err error){
	if err = db.DB.Where("id=?", id).First(story).Error; err != nil {
		return nil,err
	}
	return
}

func UpdateAStory(story *Story)(err error){
	err = db.DB.Save(story).Error
	return
}

func DeleteAStory(id string)(err error){
	err = db.DB.Where("ID=?",id).Delete(&User{}).Error
	return
}