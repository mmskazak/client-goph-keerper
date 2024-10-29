package main

import (
	"client-goph-keerper/internal/app"
	"database/sql"
	"log"
)

func main() {
	// Инициализация базы данных
	db, err := app.InitDB()
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Ошибка закрытия соединения базы данныз: %v", err)
		}
	}(db)

	//app.Start(pwdCmd, fileCmd)
}
