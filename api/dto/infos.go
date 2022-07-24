package dto

import "time"

type MessageInfo struct{
	ID         uint32    `gorm:"primary_key;auto_increment" json:"message_id"`
	Text       string    `gorm:"size:255;not null;unique" json:"text"`
	SendDate   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"send_date"`
	Sender     UserInfo  `json:"sender"`
	Receiver   UserInfo  `json:"receiver"`
}

type UserInfo struct{
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Nickname  string    `gorm:"size:255;not null;unique" json:"nickname"`
}