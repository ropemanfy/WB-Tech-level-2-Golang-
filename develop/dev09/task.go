package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("get: %v", err)
	}

	splitUrl := strings.Split(url, "/")
	fileName := splitUrl[2]

	reader, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("read: %v", err)
	}

	file, err := os.Create(fmt.Sprintf("%s.html", fileName))
	if err != nil {
		log.Fatalf("create: %v", err)
	}
	defer file.Close()

	file.WriteString(string(reader))
}
