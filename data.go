package main

type UserStatus struct {
	Paid                 bool
	PaymentType          string // "crypto" или "card"
	SelectedCountry      string
	SelectedCity         string
	SelectedOrganization string
}

var Users = map[int64]*UserStatus{}

// Доступные страны
var Countries = []string{"Словакия", "Испания", "Германия"}

// Города по странам
var Cities = map[string][]string{
	"Словакия": {"Братислава", "Нитра", "Кошице"},
	"Польша":   {"Барселона", "Мадрид"},
	"Германия": {"Берлин", "Мюнхен"},
}

// Организации по городам
var Organizations = map[string][]string{
	"Киев":    {"Организация А", "Организация Б"},
	"Львов":   {"Организация В"},
	"Одесса":  {"Организация Г"},
	"Варшава": {"Организация Д"},
	"Краков":  {"Организация Е"},
	"Берлин":  {"Организация Ж"},
	"Мюнхен":  {"Организация З"},
}
