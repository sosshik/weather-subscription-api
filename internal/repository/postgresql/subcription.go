package postgresql

import (
	"database/sql"
	"github.com/sosshik/weather-subscription-api/internal/repository"
)

func (r *PostgreSQL) CreateSubscription(sub repository.Subscription) error {
	query := `INSERT INTO subscriptions (email, city, frequency, token, is_confirmed) VALUES (:email, :city, :frequency, :token, :is_confirmed)`
	_, err := r.db.NamedExec(query, sub)
	return err
}

func (r *PostgreSQL) GetAllConfirmedSubscriptionsByFrequency(frequency string) ([]repository.Subscription, error) {
	var sub []repository.Subscription
	query := `SELECT * FROM subscriptions where is_confirmed = true and frequency = $1 order by created_at desc`
	err := r.db.Select(&sub, query, frequency)
	return sub, err
}

func (r *PostgreSQL) GetSubscriptionByToken(email string) (repository.Subscription, error) {
	var sub repository.Subscription
	query := `SELECT * FROM subscriptions where tokem = $1 order by created_at desc`
	err := r.db.Select(&sub, query, email)
	return sub, err
}

func (r *PostgreSQL) UpdateSubscription(sub repository.Subscription) error {
	query := `
		UPDATE subscriptions
		SET email = $1,
		    city = $2,
		    frequency = $3,
		    is_confirmed = $4,
		    created_at = $5
		WHERE token = $6
	`
	res, err := r.db.Exec(query,
		sub.Email,
		sub.City,
		sub.Frequency,
		sub.IsConfirmed,
		sub.CreatedAt,
		sub.Token, // used in WHERE clause
	)

	if res != nil {
		if n, rowsErr := res.RowsAffected(); n == 0 || rowsErr != nil {
			return sql.ErrNoRows
		}
	}

	return err
}

func (r *PostgreSQL) UpdateIsConfirmedSubscriptionByToken(token string) error {
	query := `
		UPDATE subscriptions
		SET is_confirmed = TRUE
		WHERE token = $1
	`
	_, err := r.db.Exec(query, token)
	return err
}
func (r *PostgreSQL) DeleteSubscriptionByToken(token string) error {
	query := `DELETE FROM subscriptions WHERE token = $1`
	res, err := r.db.Exec(query, token)

	if res != nil {
		if n, rowsErr := res.RowsAffected(); n == 0 || rowsErr != nil {
			return sql.ErrNoRows
		}
	}

	return err
}
