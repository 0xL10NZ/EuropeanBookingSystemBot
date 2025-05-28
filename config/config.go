package config

const (
	BotToken = " " // Ваш Bot Token
	AdminID  = 1234567 // Ваш Telegram ID для админки

	// Статусы
	StatusPendingPayment = "pending_payment"
	StatusPaid           = "paid"
	StatusActive         = "active"

	// Методы оплаты
	PaymentMethodCrypto = "crypto"
	PaymentMethodCard   = "card"
)
