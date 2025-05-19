package service

import (
	"github.com/google/uuid"
	"github.com/sosshik/weather-subscription-api/internal/dto"
	"github.com/sosshik/weather-subscription-api/internal/repository"
)

type SubscriptionService struct {
	repo repository.SubscriptionInterface
}

func NewSubscriptionService(repo repository.SubscriptionInterface) *SubscriptionService {
	return &SubscriptionService{repo}
}

func (s *SubscriptionService) Subscribe(sub dto.SubscribeRequestDTO) (string, error) {
	subModel := sub.ToSubscriptionModel()
	subModel.Token = uuid.New().String()
	err := s.repo.CreateSubscription(subModel)
	if err != nil {
		return "", err
	}
	return subModel.Token, nil
}

func (s *SubscriptionService) Confirm(token string) error {
	return s.repo.UpdateIsConfirmedSubscriptionByToken(token)
}

func (s *SubscriptionService) Unsubscribe(token string) error {
	return s.repo.DeleteSubscriptionByToken(token)
}
