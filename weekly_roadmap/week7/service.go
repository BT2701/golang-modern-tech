package week7

type MessageService struct {
	repo *MessageRepository
}

func NewMessageService(repo *MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) SaveMessage(message *Message) error {
	return s.repo.SaveMessage(message)
}
