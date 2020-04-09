package model

import (
	"github.com/jinzhu/gorm"
)


type User struct {
	gorm.Model
	Name      string `jason:"name"`
	Password  string `jason:"password"`
	Following int64  `json:"'following's id"`
	Follower  int64  `json:"'follower'id"`
	Avatars   Avatar `gorm:"foreignkey:AvaRefer"`
	Like      int64  `json:"like"`
	StoryId   uint64 `gorm:"association_foreignkey:Story"`
	Currency  int64
}

type Story struct{
	gorm.Model
	Story string `json:"story"`
	Image    []Img `gorm:"association_foreignkey:Img"`
}

type Img struct {
	gorm.Model
	Url string `json:"url"`
}

type Avatar struct {
    gorm.Model
	AvaRefer uint
	Url  string
}



