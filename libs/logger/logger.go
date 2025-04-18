package logger

import "log"

func Info(msg string) {
	log.Println("[INFO]", msg)
}

func Warning(msg string) {
	log.Println("[WARNING]", msg)
}

func Error(msg string) {
	log.Println("[ERROR]", msg)
}

func Debug(msg string) {
	log.Println("[DEBUG]", msg)
}
