package main

import (
	"log"
	"os"

	"github.com/DenisOzindzheDev/mgmt-file-transfer/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// Установка gin в release mode
	gin.SetMode(gin.ReleaseMode)
	// LstdFlags - ставит в лог время, Lshotfile - включает в логере параметр который указывает где сработал лог (к примеру об ошибке)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	/*
		Почему используется побитовое ИЛИ т.е ^
		const (
			Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
			Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
			Ltime                         // the time in the local time zone: 01:23:23
			Llongfile                     // full file name and line number: /a/b/c/d.go:23
			Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
			LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
			Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
			LstdFlags     = Ldate | Ltime // initial values for the standard logger
		)

		iota каждому эллементу в константе дает инкремент , константы с степенями двойки выглядят
		1 = 0001
		2 = 0010
		4 = 0100
		и.т.д
		Получается если нам нужен от логгера Ldate и Lmicroseconds то побитово это будет
		0001 + 0010 = 0011
		так передав в setFlags (0011) логер будет понимать что включены флаги 1 и 2
	*/
	// Вывод версии сервера в лог
	// os.Getenv("SERVER_VERSION") - получает значение переменной окружения SERVER_VERSION
	// если переменной нет то вернет пустую строку
	log.Printf("Starting file-transfer server: v%s", os.Getenv("SERVER_VERSION"))

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	log.Printf("Config loaded: %s", cfg.BaseURL)
	// Проверка наличия директорий и создание их, если они не существуют
	dirs := []string{cfg.DataDir, cfg.BackupDir}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatalf("Failed to create directory: %v", err)
		}
	}
}
