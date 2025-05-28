package handlers

import (
	"holdup-bot/data"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// /start — регистрация пользователя, отправка реквизитов
func HandleStart(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	user, exists := data.GetUser(chatID)
	if !exists {
		user = data.NewUser(chatID, update.Message.From.UserName)
		data.SaveUser(user)
	}

	if user.Status == data.StatusPaid {
		msg := tgbotapi.NewMessage(chatID, "Вы уже оплатили и имеете доступ. Используйте /check для проверки статуса.")
		bot.Send(msg)
		return
	}

	text := strings.Builder{}
	text.WriteString("Здравствуйте!\nДля доступа необходимо внести оплату 25$/1050грн\n\n")
	text.WriteString("Реквизиты для оплаты:\n")
	text.WriteString("Карта UA: 4102474001721279\n")
	text.WriteString("USDT TRC20: TPwkDaYXANYubNMjcA6ZhStfJwtNupji2D\n\n")
	text.WriteString("После оплаты нажмите /check для подтверждения.")

	msg := tgbotapi.NewMessage(chatID, text.String())
	bot.Send(msg)
}

// /check — проверка статуса оплаты и выдача доступа
func HandleCheck(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	user, exists := data.GetUser(chatID)
	if !exists {
		msg := tgbotapi.NewMessage(chatID, "Вы не зарегистрированы. Отправьте /start чтобы начать.")
		bot.Send(msg)
		return
	}

	switch user.Status {
	case data.StatusPaid:
		msg := tgbotapi.NewMessage(chatID, "Оплата подтверждена! Добро пожаловать.\nВыберите страну:")
		// Здесь можно добавить меню выбора страны, например:
		keyboard := tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Украина"),
				tgbotapi.NewKeyboardButton("Польша"),
			),
		)
		msg.ReplyMarkup = keyboard
		bot.Send(msg)
	case data.StatusPendingPayment:
		msg := tgbotapi.NewMessage(chatID, "Оплата не подтверждена.\nПожалуйста, оплатите по реквизитам и свяжитесь с администратором.")
		bot.Send(msg)
	}
}

// Обработка прочих сообщений
func HandleDefault(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Используйте команды /start или /check")
	bot.Send(msg)
}
