package util

import (
	"hackweek/main/db"
	"hackweek/main/model"
	"math/rand"
	"time"
)

func RandomString(n int)string {
	letters := []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM123456789")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func CreateAStory(story *model.Story)(err error) {
	err = db.DB.Create(&story).Error
	return
}

func ReadAllMyStory()(allstory []*model.Story,err error){
	if err := db.DB.Find(&allstory).Error;err!=nil{
		return  nil, err
	}
	return
}

func GETAStory(id string)(story *model.Story,err error){
	if err = db.DB.Where("id=?", id).First(new(model.Story)).Error; err != nil {
		return nil,err
	}
	return
}

func UpdateAStory(story *model.Story)(err error){
	err = db.DB.Save(story).Error;
	return
}

func DeleteAStory(id string)(err error){
	err = db.DB.Where("id=?",id).Delete(&model.Story{}).Error
	return
}