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

	if err := database.Db.Model(&model.User{}).Create(&newUser).Error; err != nil {
		panic(err)
		return newUser, err
	}
	return newUser, nil
}

func IsUserExitByName(userName string) bool {
	userExit := model.User{}
	if err := database.Db.Model(&model.User{}).Where("name=?", userName).First(userExit).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func IsUserExit(userName string, passWord string, login *model.User) error {
	if login == nil {
		return common.ErrorNullPointer
	}
	database.Db.Model(&model.User{}).Where("name=?", userName).First(login)
	if login.ID == 0 {
		return common.ErrorFullPossibility
	}
	if login.Password != passWord {
		return common.ErrorPasswordFalse
	}
	return nil
}

func GetUserByID(userID uint, userInfo *model.User) error {
	if err := database.Db.Model(&model.User{}).Where("ID=?", userID).Find(userInfo).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}

func GetUserByToken(token string, userInfo *model.User) error {
	if err := database.Db.Model(&model.User{}).Where("name=?", token).Find(userInfo).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}

func CheckIsFollow(targetID uint, userID uint) bool {
	if targetID == userID {
		return false
	}
	return IsFollowing(targetID, userID)
}
