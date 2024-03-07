package main

import (
	"fmt"
	"os"
)

const (
	msgSep = "\n"
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

func main() {
	addMessage("test.txt", "hi")
}
