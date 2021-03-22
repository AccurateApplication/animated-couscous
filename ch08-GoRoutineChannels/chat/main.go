package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string // message channel, send to

var (
	entering = make(chan client) // unsure what this is
	leaving  = make(chan client)
	messages = make(chan string) // client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all clients
	for {
		select {
		case msg := <-messages:
			// broadcast inc message to all client outgoing msg channel
			for cli := range clients {
				cli <- msg // send
			}
		case cli := <-entering:
			clients[cli] = true // client has entered and is now true

		case cli := <-leaving:
			delete(clients, cli) // what is "client" in this case and what does it do
			close(cli)
		}
	}

	//
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing ch
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You're" + who
	messages <- who + "has arrived" // send this to chat
	entering <- ch                  // clients[cli] now true

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text() // send input text to all
	}

	// this executes when input Scan is over, so C-c / C-d ?
	leaving <- ch
	messages <- who + "has left.."
	conn.Close()
}

// This prints out everything that is.. sent through channels
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
