package dto

type SendMessageRequest struct{
	Text       string    `gorm:"size:255;not null;unique" json:"text"`
	ReceiverID uint32    `gorm:"not null" json:"receiver_id"`
}