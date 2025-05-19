# Weather Subscription API

**API service for subscribing, confirming, and unsubscribing weather notifications.**

**Frontend:** [https://aquamarine-truffle-15c690.netlify.app/](https://aquamarine-truffle-15c690.netlify.app/)

---

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Technologies](#technologies)
- [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [How to use](#how-to-use)

---

## Overview

The Weather Subscription API allows users to subscribe to weather updates for cities, confirm their subscription, and unsubscribe. The service periodically sends weather updates by email based on user subscription frequency (hourly or daily).

---

## Features

- User subscription management (subscribe, confirm, unsubscribe)
- Fetching real-time weather data from an external API
- Email notifications for weather updates (hourly and daily)
- PostgreSQL database integration with migrations support
- Cron-based scheduling for sending emails
- Environment-based configuration for easy deployment

---

## Technologies

- Go (Golang)
- PostgreSQL
- sqlx (database library)
- golang-migrate (database migrations)
- robfig/cron (job scheduling)
- Logrus (logging)
- SMTP (email sending)
- dotenv (environment configuration)

---

## Getting Started

### Prerequisites

- Go 1.24+ installed
- Docker and Docker Compose installed
- SMTP email account (e.g., Gmail)
- Weather API key (OpenWeatherMap)

### How to use

1. Create `compose.env` file with this: 
```.env
CONNECTION_STRING=postgres://weather_user:weather_pass@postgres:5432/weather_db?sslmode=disable
SENDER_EMAIL=example@gmail.com
SENDER_PASS=app pass for email
POSTGRES_USER=weather_user
POSTGRES_PASSWORD=weather_pass
POSTGRES_DB=weather_db
WEATHER_API_KEY=your_api_key
SERVICE_DOMAIN=http://localhost:8090
```
2. Run docker compose:
```bash
docker compose up --build
```
Please note it may not start on the first time due to db start delay. If that is the case, stop compose and rerun command again

