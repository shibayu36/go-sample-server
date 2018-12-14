package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8000, "port to listen")
	flag.Parse()

	log.Printf("Web server starting...")
	log.Printf("Port: %d", port)
	http.HandleFunc("/echo", echoHandler)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

// logにbodyを出してレスポンスにも返す
func echoHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	body := string(b)
	log.Println(body)
	fmt.Fprint(w, body)
}
