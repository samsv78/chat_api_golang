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

func (m *Message) GetMessageInfo(db *gorm.DB) (dto.MessageInfo, error){
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