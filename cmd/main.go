package main

import (
	"dedobot/internal/db"
	"dedobot/internal/handlers"
	"dedobot/internal/repositories"
	"dedobot/internal/services"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	dbConn, err := db.InitDB()
	if err != nil {
		log.Fatal("DB error:", err)
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	repository := repositories.NewSkufRepo(dbConn)
	svc := service.NewSkufService(repository)
	botHandler := handlers.NewBotHandler(bot, svc)
	botHandler.HandleUpdates(updates)
}
