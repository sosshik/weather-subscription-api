package repository

type SubscriptionInterface interface {
	CreateSubscription(sub Subscription) error
	UpdateSubscription(sub Subscription) error
	UpdateIsConfirmedSubscriptionByToken(token string) error
	DeleteSubscriptionByToken(token string) error
	GetAllConfirmedSubscriptionsByFrequency(frequency string) ([]Subscription, error)
	GetSubscriptionByToken(email string) (Subscription, error)
}

type Repository struct {
	SubscriptionInterface
}

func NewRepository(subscription SubscriptionInterface) *Repository {
	return &Repository{
		SubscriptionInterface: subscription,
	}
}
