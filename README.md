## Build
```bash
go build -o server_asr_llm_tts cmd/main.go
```


## Run server
Make sure to set the `DASHSCOPE_API_KEY` environment variable before running the server
```bash
export DASHSCOPE_API_KEY=XXXXX
./server_asr_llm_tts
```

## Usage code
```go
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // 生产中可做跨域或 token 检查
	},
}

func ServeWs(hub *ws.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrader failed: ", err)
		return
	}

	session := ws.NewSession()

	client := &ws.Client{Hub: hub, Conn: conn, Session: session}
	client.Hub.Register <- client

	go client.ReadPump()
	go client.WritePump()
}
```
