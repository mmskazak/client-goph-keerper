package main

import (
	"bufio"
	"client-goph-keerper/internal/config"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var cfg *config.Config // глобальная переменная для конфигурации

// Инициализация корневой команды
var rootCmd = &cobra.Command{
	Use:          "keeper",
	Short:        "CLI сервер, ожидающий команды",
	SilenceUsage: true, // Отключаем справку по умолчанию
}

func main() {
	// Создаем новую конфигурацию и добавляем флаги
	cfg = config.NewConfig()
	rootCmd.PersistentFlags().StringVar(&cfg.AppUrl, "app-url", cfg.AppUrl, "URL сервера goph-keeper")
	rootCmd.PersistentFlags().StringVar((*string)(&cfg.LogLevel), "log-level", string(cfg.LogLevel), "Уровень логирования (info, debug, error)")

	// Выводим обновленные значения после выполнения команды
	fmt.Println("App URL:", cfg.AppUrl)
	fmt.Println("Log Level:", cfg.LogLevel)

	// Если аргументов нет, запускаем интерактивный режим
	interactiveMode(cfg)

}

// Функция для работы в интерактивном режиме
func interactiveMode(cfg *config.Config) {
	fmt.Println("CLI Server запущен. Введите команду или 'exit' для выхода.")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Завершаем работу, если введена команда "exit"
		if input == "" {
			continue
		}
		if input == "exit" {
			fmt.Println("Выход из программы.")
			return
		}

		// Обработка команды через Cobra
		args := strings.Split(input, " ")
		rootCmd.SetArgs(args)
		if err := rootCmd.Execute(); err != nil {
			fmt.Println("Ошибка:", err)
		}
	}
}

// Функция для печати текста с эффектом печатной машинки
func typewriterEffect(text string, delay time.Duration) {
	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(delay * time.Millisecond)
	}
	fmt.Println()
}
