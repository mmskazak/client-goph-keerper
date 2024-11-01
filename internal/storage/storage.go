package storage

import (
	"database/sql"
	"errors"
	"fmt"
)

type Storage struct {
	DataBase *sql.DB
}

func Init() (*Storage, error) {
	// Открываем базу данных SQLite (если файла нет, он будет создан)
	db, err := sql.Open("sqlite", "./keeper.db")
	if err != nil {
		return nil, fmt.Errorf("error opening db: %v", err)
	}

	// Проверяем подключение
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging db: %v", err)
	}

	// Создаем таблицу для хранения данных клиента
	// в этой таблице будет храниться
	// jwt - токен авторизованного пользователя на удаленном сервере
	// server_url - адрес удаленного сервера с секретами
	schema := `
	CREATE TABLE IF NOT EXISTS 'app_params' (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
		'key' TEXT,
		'value' TEXT
	);
	-- Вставка значения для ключа 'jwt', если такой записи нет
	INSERT OR IGNORE INTO app (key, value)
	VALUES ('jwt', '');
	       
	-- Вставка значения для ключа 'server_url', если такой записи нет
	INSERT OR IGNORE INTO app (key, value)
	VALUES ('server_url', '');
`

	// Выполнение SQL-запросов
	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil
}

// GetTokenFromDB Функция для получения токена из базы данных.
func (s *Storage) GetTokenFromDB() (string, error) {
	// Извлекаем токен из таблицы
	var token string
	query := `SELECT jwt FROM users LIMIT 1`
	err := s.db.QueryRow(query).Scan(&token)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("токен не найден в базе данных")
		}
		return "", fmt.Errorf("ошибка получения токена: %v", err)
	}

	return token, nil
}
