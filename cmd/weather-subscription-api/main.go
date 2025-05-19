package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"github.com/sosshik/weather-subscription-api/internal/emailer"
	"github.com/sosshik/weather-subscription-api/internal/weather"

	"github.com/sosshik/weather-subscription-api/internal/config"
	"github.com/sosshik/weather-subscription-api/internal/handlers"
	"github.com/sosshik/weather-subscription-api/internal/repository"
	"github.com/sosshik/weather-subscription-api/internal/repository/postgresql"
	"github.com/sosshik/weather-subscription-api/internal/service"
	"os"
)

// @title           Weather Subscription API
// @version         1.0
// @description     API service for subscribing, confirming, and unsubscribing weather notifications.
// @contact.email   support@example.com
// @host      localhost:8090
// @BasePath  /
// @schemes http https
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Warn("No .env file")
	}

	cfg := config.GetConfig()

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	db, err := sqlx.Connect("postgres", cfg.ConnectionString)
	if err != nil {
		log.Fatalln("Database connection error:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln("Database connection error:", err)
	}

	runMigrations(cfg.ConnectionString)

	repos := repository.NewRepository(postgresql.NewPostgreSQL(db))

	sender := emailer.NewEmailSender(
		cfg.SenderEmail,
		cfg.SenderPass,
		"smtp.gmail.com",
		587,
	)
	weather := weather.NewWeather(cfg.WeatherApiKey)

	services := service.NewService(repos, sender, weather)

	handler := handlers.NewHandler(services)

	c := cron.New()

	sendUpdates := func(subs []repository.Subscription) {
		cityToEmailsMap := make(map[string][]string)
		for _, sub := range subs {
			cityToEmailsMap[sub.City] = append(cityToEmailsMap[sub.City], sub.Email)
		}

		for city, emails := range cityToEmailsMap {
			w, weatherErr := services.GetWeather(city)
			if weatherErr != nil {
				log.Errorf("Unable to send weather update: %v", weatherErr)
				continue
			}

			body := fmt.Sprintf("Current Temperature is %d\nCurrent Humidity is %d\nWeather descreiption: %s", w.Temperature, w.Humidity, w.Description)
			sendErr := sender.Send(
				emails,
				fmt.Sprintf("Weather update for %s", city),
				body,
			)
			if sendErr != nil {
				log.Errorf("Unable to send updates email: %v", sendErr)
			}
		}
	}

	c.AddFunc("28 * * * *", func() {
		subs, subsErr := repos.GetAllConfirmedSubscriptionsByFrequency("hourly")
		if subsErr != nil {
			log.Error(err)
			return
		}

		sendUpdates(subs)

	})

	c.AddFunc("0 6 * * *", func() {
		subs, subsErr := repos.GetAllConfirmedSubscriptionsByFrequency("daily")
		if subsErr != nil {
			log.Error(err)
			return
		}

		sendUpdates(subs)

	})
	c.Start()

	srv := handler.InitRoutes()
	log.Fatal(srv.Start(":8090"))
}

func runMigrations(databaseURL string) {
	m, err := migrate.New(
		"file://migrations",
		databaseURL,
	)
	if err != nil {
		log.Fatalf("Migration init error: %v", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Info("Migrations applied successfully.")
}
