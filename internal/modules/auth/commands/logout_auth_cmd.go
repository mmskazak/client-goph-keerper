package commands

import (
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out the current user",
	RunE: func(cmd *cobra.Command, args []string) error {
		removeTokenFromDB()
		fmt.Println("You have logged out")
		return nil
	},
}

func InitLogoutCmd() *cobra.Command {
	// Можно добавить флаг для токена, если требуется
	return logoutCmd
}

func removeTokenFromDB() error {
	// Подключаемся к базе данных с драйвером glebarez/sqlite
	db, err := sql.Open("sqlite", "gophkeeper.db")
	if err != nil {
		return fmt.Errorf("ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	// Вставляем токен в таблицу
	insertQuery := `DELETE FROM users WHERE 1`
	if _, err := db.Exec(insertQuery); err != nil {
		return fmt.Errorf("ошибка удаленния jwt токена: %v", err)
	}
	return nil
}
