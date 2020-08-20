package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"
)

type Message struct {
	Name string `json:"name"`
	Password string `json:"password"`
}

// define a channel
var channel = make(chan Message)

var clients = make(map[*websocket.Conn]bool)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade", err)
		return
	}

	defer func() {
		log.Println("关闭连接")
		c.Close()
	}()

	// Register new Client
	clients[c] = true

	for {
		var message Message
		err := c.ReadJSON(&message)
		if err != nil {
			log.Println("read:", err)
			break
		}

		channel <- message
	}
}

func handleMessage() {
	log.Println("handleMessage")

	for {
		message := <-channel
		for conn, value := range clients {
			log.Println("index: %s", value)
			err := conn.WriteJSON(message)

			if err != nil {
				log.Fatal("failed")
			}
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://" + r.Host+"/echo")
}

func main()  {
	flag.Parse()
	log.SetFlags(0)

	// deal with data
	go handleMessage()

	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;
    var print = function(message) {
        var d = document.createElement("div");
        d.textContent = message;
        output.appendChild(d);
    };
    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };
    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };
    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };
});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))
