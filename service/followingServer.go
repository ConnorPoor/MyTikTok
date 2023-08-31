package service

import (
	"errors"
	"github.com/ConnorPoor/MyTikTok/database"
	"github.com/ConnorPoor/MyTikTok/model"
	"gorm.io/gorm"
)

func IsFollowing(HostId uint, GuestId uint) bool {
	var relationExist = &model.Following{}
	if err := database.Db.Model(&model.Following{}).Where("host_id=? AND guest_id=?", HostId, GuestId).First(relationExist).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
