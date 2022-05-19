package main

import (
	"github.com/rheola/ozon-bot/internal/app/commands"
	"github.com/rheola/ozon-bot/internal/service/product"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	productService := product.NewService()
	commander := commands.NewCommander(bot, productService)

	uc := tgbotapi.UpdateConfig{
		Timeout: 60,
	}
	updates := bot.GetUpdatesChan(uc)
	for update := range updates {
		commander.HandleUpdate(update)
	}
}
