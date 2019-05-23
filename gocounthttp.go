package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var Log *log.Logger

type MyHandler struct {
	sync.Mutex
	count int
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Visitor count: %d. \n", h.count)
	h.Lock()
	h.count++
	h.Unlock()
	Log.Printf("Visitors: %d \n", h.count)
	if _, err := os.Stat("/503"); err == nil {
		w.WriteHeader(503)
	}
	time.Sleep(5 * time.Millisecond)
	h.Lock()
	h.count--
	h.Unlock()
}

const addr = "0.0.0.0:8080"

func main() {
	errorlog, err := os.OpenFile("/gocount.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	Log = log.New(errorlog, "applog: ", log.Lshortfile|log.LstdFlags)
	mux := http.NewServeMux()
	handler := &MyHandler{}
	mux.Handle("/", handler)
	Log.Printf("Now listening on %s...\n", addr)
	log.Printf("Now listening on %s...\n", addr)
	server := http.Server{Handler: mux, Addr: addr}
	log.Fatal(server.ListenAndServe())
}
