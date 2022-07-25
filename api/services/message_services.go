package services

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	amqp "github.com/rabbitmq/amqp091-go"
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

func SignalViaRabbit(senderID uint32, receiverID uint32)(error){
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil{
		return err
	}
	// failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil{
		return err
	}
	// failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_direct", // name
		"direct",      // type
		false,         // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil{
		return err
	}
	// failOnError(err, "Failed to declare an exchange")

	// send message 
	strSenderID := fmt.Sprint(senderID)
	strReceiverID := fmt.Sprint(receiverID)
	body := strSenderID + "," + strReceiverID
	err = ch.Publish(
		"logs_direct", // exchange
		strReceiverID, // routing key
		false,		   // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil{
		return err
	}
	// failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
	return nil
}