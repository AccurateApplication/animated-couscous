package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler) // each request calls handler?
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path) //
	for i := 1; i < 10; i++ {
		fmt.Fprintf(w, "%d testing\n", i)
		fmt.Fprintln(w, "chong", i)
	}
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "Remote addr = %q\n", r.RemoteAddr)
	fmt.Fprintf(w, "host = %q\n", r.Host)
	//fmt.Fprintf(w, "host = %q\n", r.Form)
	err := r.ParseForm()
	if err != nil {
		log.Print(err)
	}
	//for _, x := range r.Form {
	//	fmt.Fprintf(w, "r form: %q\n", x)
	//}
	for i, val := range r.Form {
		fmt.Fprintf(w, "r form i: %q\t form val: %q\n", i, val)
	}
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "count%d\n", count)
	mu.Unlock()
}
