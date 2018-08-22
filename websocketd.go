// Borrowed heavily from https://gist.github.com/tmichel/7390690
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type msg struct {
	Num int
}

func main() {
	http.HandleFunc("/ws_echo", wsHandler)
	http.HandleFunc("/ws_broadcast", wsHandler)
	http.HandleFunc("/", rootHandler)

	panic(http.ListenAndServe(":8010", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Println("Could not open file.", err)
	}
	fmt.Fprintf(w, "%s", content)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Origin") != "http://"+r.Host {
		fmt.Println("Origin not allowed")
		fmt.Printf("Expected:http://%s Actual:%s", r.Host, r.Header.Get("Origin"))
		http.Error(w, "Origin not allowed", 403)
		return
	}
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	if r.URL.Path == "/ws_echo" {
		go echo(conn)
	} else if r.URL.Path == "/ws_broadcast" {
		go broadcast(conn)
	}
}

func echo(conn *websocket.Conn) {
	for {
		m := msg{}

		err := conn.ReadJSON(&m)
		if err != nil {
			fmt.Println("Error reading json.", err)
		}

		fmt.Printf("Got message: %#v\n", m)

		if err = conn.WriteJSON(m); err != nil {
			fmt.Println(err)
		}
	}
}

func broadcast(conn *websocket.Conn) {
	m := msg{}
	m.Num = 0
	for {
		m.Num++
		var err error
		if err = conn.WriteJSON(m); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(m.Num)
		}
		time.Sleep(time.Second)
	}
}
