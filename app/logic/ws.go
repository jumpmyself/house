package logic

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync" // 用于互斥锁
)

var (
	clients   = make(map[*websocket.Conn]bool) // 存储所有连接的客户端
	broadcast = make(chan Message)             // 广播消息通道
	mutex     sync.Mutex                       // 互斥锁
)

// Message 结构用于处理接收到的消息
type Message struct {
	Author string `json:"author"`
	Text   string `json:"text"`
}

// 升级HTTP连接为WebSocket连接
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// StartServer 启动WebSocket服务器
func StartServer(port string) {
	// WebSocket路由
	http.HandleFunc("/ws", handleWebSocket)

	// 启动广播监听器
	go handleMessages()

	// 启动服务器
	log.Printf("WebSocket server started on port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// 处理WebSocket连接
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		mutex.Lock()
		delete(clients, conn)
		mutex.Unlock()
		conn.Close()
	}()

	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			break
		}

		broadcast <- msg
	}
}

// 启动广播监听器
func handleMessages() {
	for {
		msg := <-broadcast

		mutex.Lock()
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Println(err)
				client.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()
	}
}
