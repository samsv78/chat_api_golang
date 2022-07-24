package services

import (
	"github.com/jinzhu/gorm"
	"github.com/samsv78/chat_api_golang/api/dto"
	"github.com/samsv78/chat_api_golang/api/helpers"
	"github.com/samsv78/chat_api_golang/api/models"
)

func GetChatRoomsInfo(db *gorm.DB, userID uint32)([]dto.ChatRoomInfo, error){
	chatrooms := []dto.ChatRoomInfo{}
	messages, err := models.GetMessagesByUserID(db, userID)
	if err != nil{
		return []dto.ChatRoomInfo{}, err
	}
	otherUserIDs := FindOtherUserIds(messages, userID)
	for _, otherUserID := range otherUserIDs {
		chatroom, err := GetChatRoomInfo(db, userID, otherUserID)
		if err != nil{
			return []dto.ChatRoomInfo{}, err
		}
		chatrooms = append(chatrooms, chatroom)
	}
	return chatrooms, nil
}

func FindOtherUserIds(messages []models.Message, userID uint32)([]uint32){
	otherUserIDs := []uint32{}
	for _, message := range messages {
		if userID == message.SenderID{
			otherUserIDs = append(otherUserIDs, message.ReceiverID)
		}else{
			otherUserIDs = append(otherUserIDs, message.SenderID)
		}
	}
	otherUserIDs = helpers.RemoveDuplicate(otherUserIDs)
	return otherUserIDs

}

func GetChatRoomInfo(db *gorm.DB, userID uint32, otherUserID uint32)(dto.ChatRoomInfo, error){
	user, err := GetUserInfo(db, userID)
	if err != nil{
		return dto.ChatRoomInfo{}, err
	}

	otherUser, err := GetUserInfo(db, otherUserID)
	if err != nil{
		return dto.ChatRoomInfo{}, err
	}
	messages, err := models.GetMessagesOfTwoUsers(db, userID, otherUserID)
	if err != nil{
		return dto.ChatRoomInfo{}, err
	}
	messagesInfo := GetMessagesInfo(messages, user, otherUser)
	chatRoomInfo := dto.ChatRoomInfo{
		User: user,
		OtherUser: otherUser,
		Messages: messagesInfo,
	}
	return chatRoomInfo, nil
}

func GetMessagesInfo(messages []models.Message, user1 dto.UserInfo, user2 dto.UserInfo) ([]dto.MessageInfo){
	messagesInfo := []dto.MessageInfo{}
	for _, message := range messages {
		messageInfo := dto.MessageInfo{
			ID: message.ID,
			Text: message.Text,
			SendDate: message.SendDate,
		}
		if message.SenderID == user1.ID{
			messageInfo.Sender = user1
			messageInfo.Receiver = user2
		}else{
			messageInfo.Sender = user2
			messageInfo.Receiver = user1
		}
		messagesInfo = append(messagesInfo, messageInfo)
	}
	return messagesInfo
}

func GetMessageInfo(db *gorm.DB, m models.Message) (dto.MessageInfo, error){
	senderInfo, err := GetUserInfo(db, m.SenderID)
	if err != nil{
		return dto.MessageInfo{}, err
	}
	receiverInfo, err := GetUserInfo(db, m.ReceiverID)
	if err != nil{
		return dto.MessageInfo{}, err
	}
	messageInfo := dto.MessageInfo{
		ID: m.ID,
		Text: m.Text,
		SendDate: m.SendDate,
		Sender: senderInfo,
		Receiver: receiverInfo,
	}
	return messageInfo, nil
}

