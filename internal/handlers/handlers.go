// internal/handlers/handlers.go
package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// GetHTML возвращает HTML форму
func GetHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}

func UploadHTML(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Unable to get file from form", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read file content", http.StatusInternalServerError)
		return
	}
	originalContent := strings.TrimSpace(string(fileContent))
	convertedContent := service.TextConverter(string(fileContent))
	if convertedContent == "" {
		http.Error(w, "Unable empty string", http.StatusInternalServerError)
		return
	}

	timestamp := time.Now().UTC().Format("2006-01-02_15-04-05")
	ext := filepath.Ext(handler.Filename)
	outputFilename := fmt.Sprintf("%s%s", timestamp, ext)

	err = os.WriteFile(outputFilename, []byte(convertedContent), 0755)
	if err != nil {
		http.Error(w, "Unable to write in output file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(fmt.Sprintf(
		"Original file: %s\nSaved as: %s\n\nOriginalText: %s\n\nResult:\n\n%s",
		handler.Filename, outputFilename, originalContent, convertedContent,
	)))
	if err != nil {
		http.Error(w, "Unable to write on server", http.StatusInternalServerError)
		return
	}
}
