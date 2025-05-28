package handlers

import (
	"fmt"
	"holdup-bot/config"
	"holdup-bot/data"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Админ подтверждает оплату командой: /confirm <user_id>
func HandleAdmin(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	if chatID != config.AdminID {
		msg := tgbotapi.NewMessage(chatID, "Вы не администратор.")
		bot.Send(msg)
		return
	}

	args := strings.Fields(update.Message.Text)
	if len(args) < 2 || args[0] != "/confirm" {
		msg := tgbotapi.NewMessage(chatID, "Используйте: /confirm <user_id>")
		bot.Send(msg)
		return
	}

	userID, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		msg := tgbotapi.NewMessage(chatID, "Неверный ID пользователя.")
		bot.Send(msg)
		return
	}

	user, exists := data.GetUser(userID)
	if !exists {
		msg := tgbotapi.NewMessage(chatID, "Пользователь не найден.")
		bot.Send(msg)
		return
	}

	user.Status = data.StatusPaid
	data.SaveUser(user)

	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Пользователь %d подтверждён как оплативший.", userID))
	bot.Send(msg)

	// Уведомление пользователю
	notify := tgbotapi.NewMessage(userID, "Ваша оплата подтверждена администратором. Добро пожаловать!")
	bot.Send(notify)
}
