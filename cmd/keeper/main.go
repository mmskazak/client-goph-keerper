package main

import (
	"bufio"
	"client-goph-keerper/internal/modules/auth"
	"client-goph-keerper/internal/modules/file"
	"client-goph-keerper/internal/modules/pwd"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// Инициализация корневой команды
var rootCmd = &cobra.Command{
	Use:          "keeper",
	Short:        "CLI сервер, ожидающий команды",
	SilenceUsage: true, // Отключаем справку по умолчанию
}

func main() {
	// Добавляем команды к корневой команде
	rootCmd.AddCommand(auth.InitAuthCmd())
	rootCmd.AddCommand(pwd.InitPwdCmd())
	rootCmd.AddCommand(file.InitFileCmd())

	interactiveMode()
}

// Функция для работы в интерактивном режиме
func interactiveMode() {
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
