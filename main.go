package main

import (
	"log"

	"holdup-bot/config"
	"holdup-bot/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				handlers.HandleStart(bot, update)
			case "check":
				handlers.HandleCheck(bot, update)
			case "confirm":
				handlers.HandleAdmin(bot, update)
			default:
				handlers.HandleDefault(bot, update)
			}
		} else {
			handlers.HandleDefault(bot, update)
		}
	}
}
