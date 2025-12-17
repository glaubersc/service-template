package logger

import (
	"log"
)

func Info(msg string, fields map[string]any) {
	log.Println(appendFields("INFO", msg, fields))
}

func Error(msg string, fields map[string]any) {
	log.Println(appendFields("ERROR", msg, fields))
}

func appendFields(level, msg string, fields map[string]any) map[string]any {
	fields["level"] = level
	fields["message"] = msg
	return fields
}
