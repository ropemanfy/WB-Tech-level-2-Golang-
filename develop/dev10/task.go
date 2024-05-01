package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	var timeout time.Duration
	network := "tcp"

	flag.DurationVar(&timeout, "timeout", 10*time.Second, "connection timeout")
	flag.Parse()

	host := flag.Arg(len(flag.Args()) - 2)
	port := flag.Arg(len(flag.Args()) - 1)

	addr := fmt.Sprintf("%s:%s", host, port)

	conn, err := net.DialTimeout(network, addr, timeout)
	if err != nil {
		log.Fatalf("connect: %v", err)
	}
	defer conn.Close()

	go write(&conn)

	read()
}

func write(conn *net.Conn) {
	text, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
	if err == io.EOF {
		log.Println("Stop signal received. Connection closing")
		return
	}
	_, err = (*conn).Write(text)
	if err != nil {
		log.Fatalf("write: %v", err)
	}
}

func read() {
	for {
		text, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Println("connection closed")
			os.Exit(0)
		}
		fmt.Printf("socket message recieved: %v\n", text)
	}
}
