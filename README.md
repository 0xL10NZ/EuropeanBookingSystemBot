# EuropeanBookingSystemBot

A Telegram bot for managing reservations in various organizations across Europe. Access is granted only after the user completes a payment (via cryptocurrency or Ukrainian bank card) and receives manual admin confirmation.

## 🌍 Overview

**EuropeanBookingSystemBot** is a secure and interactive Telegram bot designed to help users book spots in different organizations across Europe. The bot restricts access until payment is confirmed by an admin.

## ✅ Features

- 💸 Payment via crypto or Ukrainian bank card
- 🔐 Access control with admin confirmation
- 🌍 Step-by-step selection of country, city, and organization
- 💬 Intuitive Telegram interaction
- 🧠 Simple in-memory user state management
- 🛠️ Modular structure for easy updates and scalability

## 📁 Project Structure

EuropeanBookingSystemBot/
├── main.go               # Entry point of the application
├── go.mod                # Go module definition
├── go.sum                # Go dependencies checksum
├── config/
│   └── config.go         # Bot configuration (token, admin ID, etc.)
├── handlers/
│   ├── user.go           # Handles user commands and flow
│   └── admin.go          # Admin tools and payment confirmations
├── data/
│   └── storage.go        # In-memory storage for user data
└── assets/
    └── branding.txt      # Bot welcome message and branding text

## ⚙️ Installation

### 1. Clone the repository
```bash```
git clone https://github.com/yourusername/EuropeanBookingSystemBot.git
cd EuropeanBookingSystemBot

2. Install Go
Make sure you have Go 1.20+ installed: https://go.dev/doc/install

3. Install dependencies

go mod tidy

4. Configure the bot

Open config/config.go and set your:
Telegram Bot Token
Admin Telegram ID

🚀 Running the Bot
Simply run:

go run main.go

The bot will start and listen for updates.

💡 How It Works
User sends /start
Chooses payment method: card or crypto
Receives payment details
Sends /check after payment
Admin is notified and manually confirms
User is granted access to booking steps:
Choose country
Choose city
Choose organization

🤝 Contributing
Contributions are welcome!

Please open an issue or pull request with your improvements or bugfixes.

📜 License
MIT License © 2025 0xL10NZ

📬 Contact
Telegram: @OxL10NZE
