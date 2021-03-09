package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

var tzFlag = flag.String("tz", "Local", "Timezone, for example Asia/Shanghai ")
var PortFlag = flag.Int("port", 8000, "Port to serve ")

func main() {
	flag.Parse()
	fmt.Println(*PortFlag)

	//Tz, err := TimeIn(time.Now(), *tzFlag)
	//if err != nil {
	//	fmt.Println(err)
	//}
	httpListen(*PortFlag)
}

func httpListen(port int) {
	portStr := strconv.Itoa(port)
	address := "localhost:" + portStr
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	for {
		con, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleCon(con)
	}
}

func handleCon(c net.Conn) {
	defer c.Close()
	//tz, err := TimeIn(time.Now(), "Asia/Shanghai")
	//if err != nil {
	//tzFormat
	for {
		time123, err := TimeIn(time.Now(), *tzFlag)
		if err != nil {
			log.Fatal(err)
		}
		timeString := time123.Format("15:04:05") + "\n"
		_, err = io.WriteString(c, timeString)
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func TimeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}
