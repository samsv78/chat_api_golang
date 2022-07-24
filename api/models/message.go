package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/samsv78/chat_api_golang/api/dto"
)

type Message struct{
	ID         uint32    `gorm:"primary_key;auto_increment" json:"message_id"`
	Text       string    `gorm:"size:255;not null;unique" json:"text"`
	SendDate   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"send_date"`
	SenderID   uint32    `gorm:"not null" json:"sender_id"`
	ReceiverID uint32    `gorm:"not null" json:"receiver_id"`
}

func (m *Message) Ctor(request dto.SendMessageRequest, senderID uint32) {
	m.Text = request.Text
	m.SendDate = time.Now()
	m.SenderID = senderID
	m.ReceiverID = request.ReceiverID
}

func (m *Message) SaveMessage(db *gorm.DB) (error) {

	var err error
	err = db.Debug().Create(&m).Error
	if err != nil {
		return err
	}
	return nil
}

func GetMessagesOfTwoUsers(db *gorm.DB, userID1 uint32, userID2 uint32)([]Message, error){
	messages := []Message{}
	if result := db.Where("sender_id = ? AND receiver_id = ? OR sender_id = ? AND receiver_id = ?", userID1, userID2, userID2, userID1).Find(&messages); result.Error != nil {
		return messages, result.Error
	}
	return messages, nil
}

func GetMessagesByUserID(db *gorm.DB, userID uint32)([]Message, error){
	messages := []Message{}
	if result := db.Select("sender_id, receiver_id").Where("sender_id = ? OR receiver_id = ?", userID, userID).Find(&messages); result.Error != nil {
		return messages, result.Error
	}
	return messages, nil
}