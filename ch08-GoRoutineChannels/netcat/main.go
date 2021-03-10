package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	fmt.Printf("con: %T\n", conn)
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{}) // why struct
	go func() {
		io.Copy(os.Stdout, conn) // ignoring err
		log.Println("Done")
		done <- struct{}{} // send to main routine
	}()
	mustCopy(conn, os.Stdin) // write from stdin to conn
	conn.Close()
	<-done
}

// io write from io reader
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
