package main

import (
	"client-goph-keeper/internal/app"
	"client-goph-keeper/internal/modules/file"
	"client-goph-keeper/internal/modules/pwd"
)

func main() {
	//Командная оболочка для паролей
	pwdCmd := pwd.Init()
	//Командная оболочка для файлов
	fileCmd := file.Init()
	app.Start(pwdCmd, fileCmd)
}
