package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

const (
	msgSep       = "\n"
	htmlFilename = "index.html"
)

func addMessage(filename string, message string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(fmt.Sprintf("Failed to open file " + filename))
	}

	fileInfo, err := file.Stat()

	if err != nil {
		panic(fmt.Sprintf("Failed to get stat for file " + filename))
	}

	if fileInfo.Size() != 0 {
		if _, err := file.WriteString(msgSep); err != nil {
			panic(fmt.Sprintf("Failed to write message separator in file " + filename))
		}
	}

	if _, err := file.WriteString(message); err != nil {
		panic(fmt.Sprintf("Failed to write message " + message + " in file " + filename))
	}

	defer file.Close()
}

func readMessages(filename string) string {
	b, err := os.ReadFile(filename)

	if err != nil {
		panic(fmt.Sprintf("Failed to read file " + filename))
	}

	return string(b)
}

func handler(w http.ResponseWriter, r *http.Request) {
	messages := map[string]interface{}{"messages": readMessages("messages.txt")}

	t, _ := template.ParseFiles(htmlFilename)
	t.Execute(w, messages)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
