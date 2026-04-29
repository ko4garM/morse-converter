package service

import (
	"log"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func TextConverter(text string) string {
	trimmedWord := strings.TrimSpace(text)
	if len(trimmedWord) == 0 {
		log.Println("Ошибка, передана пустая строка")
		return ""
	}
	isMorse := true
	for _, ch := range trimmedWord {
		if ch != '.' && ch != '-' && ch != ' ' {
			isMorse = false
			break
		}
	}
	if isMorse {
		return morse.ToText(trimmedWord)
	}
	return morse.ToMorse(trimmedWord)
}
