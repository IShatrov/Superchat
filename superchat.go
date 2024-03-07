package main

import (
	"log"
	"os"
)

func addMessage(filename string, message string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	fileInfo, err := file.Stat()

	if err != nil {
		log.Fatal(err)
	}

	if fileInfo.Size() != 0 {
		if _, err := file.Write([]byte("\n")); err != nil {
			file.Close()
			log.Fatal(err)
		}
	}

	if _, err := file.Write([]byte(message)); err != nil {
		file.Close()
		log.Fatal(err)
	}

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	addMessage("test.txt", "hi")
}
