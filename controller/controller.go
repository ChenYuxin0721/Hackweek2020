package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"hackweek/main/db"
	"hackweek/main/dto"
	"hackweek/main/model"
	"hackweek/main/response"
	"hackweek/main/util"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	DB := db.GetDB()
	name := c.Query("name")
	password := c.Query("password")
	var user model.User
	DB.Where("name=?", name).First(&user)
	if user.ID != 0 {
		response.Response(c, http.StatusOK, 400, nil, "用户名已存在")
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	if len(password) < 6 {
		response.Response(c, http.StatusOK, 400, nil, "密码不能少于6位")
		return
	} //创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusOK, 400, nil, "加密错误")
		return
	}
	newUser := model.User{
		Model:     gorm.Model{},
		Name:      name,
		Password:  string(hasedPassword),
		Following: 0,
		Follower:  0,
		Like:      0,
		StoryId:   0,
		Currency:  0,
	}

	DB.Create(&newUser)

	response.Success(c, nil, "注册成功")
}

func Login(c *gin.Context) {
	DB := db.GetDB()
	//获取参数
	name := c.Query("name")
	password := c.Query("password")
	//数据验证
	if len(name) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "用户名不能为空"})
		return
	}
	if len(password) < 6 {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "密码不能少于6位"})
		return
	}
	//判断用户是否存在
	var user model.User
	DB.Where("name=?", name).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "用户不存在"})
		return
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "密码错误"})
		return
	}
	//发放token
	token, err := db.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    400,
			"message": "系统异常"})
		log.Printf("token generate error: %v", err)
		return
	}
	//返回结果
	response.Success(c, gin.H{"token": token}, "登录成功")
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"user": dto.ToUserDTO(user.(model.User))}})
}

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func CreateAStory(c *gin.Context) {
	//获取参数
	name := c.Query("name")
	Text := c.Query("text")
	Title := c.Query("title")
	Tag := c.Query("tag")
	Imagurl := c.Query("Imagurl")
	//数据验证
	if len(Text) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": "400", "message": "内容不能为空"})
		return
	}
	if len(Title) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": "400", "message": "内容不能为空"})
		return
	}
	if len(Tag) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": "400", "message": "内容不能为空"})
		return
	}

	newStory := model.Story{
		Model:   gorm.Model{},
		Name:    name,
		Title:   Title,
		Text:    Text,
		Tag:     Tag,
		Imagurl: Imagurl,
	}
	db.DB.Create(&newStory)

	response.Success(c, gin.H{
		"id":      newStory.ID,
		"name":    name,
		"title":   Title,
		"text":    Text,
		"tag":     Tag,
		"imagurl": Imagurl,
	},
		" 发表成功")

}

func ReadAllMyStory(c *gin.Context) {
	//查询我的故事的所有数据
	name := c.Query("name")
	var allstory []model.Story
	db.DB.Where(&model.Story{Name: name}).First(&allstory)
	db.DB.Table("Story").Where("name = ?", "name ").Find(&allstory)
	response.Success(c, nil, "删除成功")
}

func UpdateAStory(c *gin.Context) {
	id, ok := c.Params.Get("Id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效ID"})
		return
	}
	story, err := util.GETAStory(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&story)
	if err = util.UpdateAStory(story); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "保存失败",})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "更新成功",
			"data":    story,
		})
	}
}

func DeleteAStory(c *gin.Context) {
	DB := db.GetDB()
	id := c.Query("id")
	story := db.DB.Table("Story").Where("id = ?", id).First(&id)
	DB.Delete(&story)
	response.Success(c, nil, "删除成功")

}

func DeleteUser(c *gin.Context) {

	DB := db.GetDB()
	id := c.Query("id")
	user := db.DB.Table("User").Where("id = ?", id).First(&id)
	DB.Delete(&user)
	response.Success(c, nil, "删除成功")
}
