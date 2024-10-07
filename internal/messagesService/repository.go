package messagesService

import "gorm.io/gorm"

type MessageRepository interface {
	CreateMessage(message Message) (Message, error)

	GetAllMessages() ([]Message, error)

	UpdateMessageByID(id uint, message Message) (Message, error)

	DeleteMessageByID(id uint) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *messageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) CreateMessage(message Message) (Message, error) {
	result := r.db.Create(&message)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return message, nil
}

func (r *messageRepository) GetAllMessages() ([]Message, error) {
	var messages []Message
	err := r.db.Find(&messages).Error
	return messages, err
}

func (r *messageRepository) UpdateMessageByID(id uint, message Message) (Message, error) {

	existingMessage := r.db.First(&message, id)
	if existingMessage.Error != nil {
		return Message{}, existingMessage.Error
	}

	updatedMessage := r.db.Model(&existingMessage).Updates(message)
	if updatedMessage.Error != nil {
		return Message{}, updatedMessage.Error
	}

	return message, nil
}

func (r *messageRepository) DeleteMessageByID(id uint) error {
	message := Message{}
	result := r.db.Delete(&message, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
