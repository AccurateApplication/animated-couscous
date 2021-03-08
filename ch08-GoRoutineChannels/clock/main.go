package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8000") // create tcp listen on 8000
	if err != nil {
		log.Fatal(err)
	}

	for { // while
		fmt.Println("start of while")
		netCat()
		conn, err := listen.Accept() // conn = accept LH:8000
		if err != nil {
			log.Print(err) // conn failed/aborted
			continue       // continue for/while loop

		}
		go handleConn(conn) // one at a time
		go netCat()
	}

}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n")) // why the "_"
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
		//ncPrint := netCat("localhost:8000")
	}
}

// This seems to print when i runt NC on localhost..
func netCat() {
	fmt.Println("netcat")
	con, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()
	ioCopy(os.Stdout, con)
}
func ioCopy(dst io.Writer, src io.Reader) {
	fmt.Println("iocopy")
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
