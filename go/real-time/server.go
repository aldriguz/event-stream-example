package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

var addr *string = flag.String("addr", ":3000", "address")

func main() {
	flag.Parse()

	http.HandleFunc("/", home)
	http.HandleFunc("/events", events)
	http.ListenAndServe(*addr, nil)
}

func events(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")

	tokens := []string{"example", "of", "real", "time", "response", "from", "server"}

	for _, token := range tokens {
		content := fmt.Sprintf("data: %s\n\n", string(token))
		w.Write([]byte(content))
		w.(http.Flusher).Flush()

		//this is intentional
		time.Sleep(time.Millisecond * 420)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("server is up!"))
}
