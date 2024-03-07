package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	msgSep      = "\n"
	msgFilename = "messages.txt"

	htmlFilename = "index.html"
	htmlBody     = "body"
	htmlMessages = "messages"

	port = ":8080"

	chatPath = "/chat"
	sendPath = "/send"

	timeFormat = "2006-01-02 15:04:05"
)

func addMessage(filename string, message string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	defer file.Close()

	if err != nil {
		return err
	}

	fileInfo, err := file.Stat()

	if err != nil {
		return err
	}

	if fileInfo.Size() != 0 {
		if _, err := file.WriteString(msgSep); err != nil {
			return err
		}
	}

	if _, err := file.WriteString(message); err != nil {
		return err
	}

	return nil
}

func readMessages(filename string) (string, error) {
	b, err := os.ReadFile(filename)

	if err != nil {
		return "", nil
	}

	return string(b), nil
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	text, err := readMessages(msgFilename)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	messages := map[string]interface{}{htmlMessages: text}

	t, err := template.ParseFiles(htmlFilename)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, messages)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue(htmlBody)

	t := time.Now()

	err := addMessage(msgFilename, t.Format(timeFormat)+" "+body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, chatPath, http.StatusFound)
}

func main() {
	http.HandleFunc(chatPath, chatHandler)
	http.HandleFunc(sendPath, saveHandler)
	log.Fatal(http.ListenAndServe(port, nil))
}
