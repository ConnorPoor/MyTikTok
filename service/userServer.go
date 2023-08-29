package service

import (
	"errors"
	"github.com/ConnorPoor/MyTikTok/common"
	"github.com/ConnorPoor/MyTikTok/database"
	"github.com/ConnorPoor/MyTikTok/model"
	"gorm.io/gorm"
)

func CreateUser(userName string, passWord string) (model.User, error) {

	var newUser = model.User{
		Name:           userName,
		Password:       passWord,
		FollowCount:    0,
		FollowerCount:  0,
		FavoriteCount:  0,
		TotalFavorited: 0,
	}

	if IsUserExitByName(userName) {
		return newUser, common.ErrorUserExit
	}

	if err := database.Db.Model(model.User{}).Create(&newUser).Error; err != nil {
		panic(err)
		return newUser, err
	}
	return newUser, nil
}

func IsUserExitByName(userName string) bool {
	userExit := model.User{}
	if err := database.Db.Model(model.User{}).Where("name=?", userName).First(userExit).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
