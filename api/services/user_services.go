package services

import (
	"github.com/jinzhu/gorm"
	"github.com/samsv78/chat_api_golang/api/dto"
	"github.com/samsv78/chat_api_golang/api/models"
)

func GetUserInfo(db *gorm.DB, uid uint32) (dto.UserInfo, error){
	user := models.User{}
	u, err := user.FindUserByID(db, uid)
	user = *u
	if err != nil{
		return dto.UserInfo{}, err
	}
	userInfo := dto.UserInfo{
		ID: user.ID,
		Nickname: user.Nickname,
	}
	return userInfo, err
}