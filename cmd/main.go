package main

import (
	"client-goph-keeper/internal/app"
	"log"
)

func main() {
	// Инициализация базы данных
	db, err := app.InitDB()
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}
	defer db.Close()

	//app.Start(pwdCmd, fileCmd)
}
