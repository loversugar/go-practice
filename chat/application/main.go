package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var add = flag.String("addr", "localhost:8080", "http service address")
var upgrader = websocket.Upgrader{}

func main()  {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
}

func echo(w http.ResponseWriter, r *http.Request) {

}
