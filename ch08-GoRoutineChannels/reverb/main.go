package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {

	listen, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Print(err)
			continue

		}
		go handleCon(conn)
	}

}

func echo(c net.Conn, word string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(word))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", word)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(word))
}

func handleCon(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second) // each time get input, goroutine
	}
	c.Close()
}
