package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	go helloGoodbye() // this prints
	start := time.Now()
	ch := make(chan string) // what is this
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // go routine, hmm
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // recieving from channel ch (?)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	//go helloGoodbye() // this wont print
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // sending to channel ch ?
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("err reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%2fs %7d %s", secs, nbytes, url) // confusing
	// also look up this confusing Sprintf formatting
}

func helloGoodbye() { // go routine
	fmt.Println("hello goodbye")
	a := make(chan int) // go channel
	fmt.Printf("type:%T\t as string: %s\t as int: %d\t\n", a, a, a)
	fmt.Println(a)
	a <- 5867234
	fmt.Println("redicercted to channel, i think")
	fmt.Printf("type:%T\t as string: %s\t as int: %d\t\n", a, a, a)
}
