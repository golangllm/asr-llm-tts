## Build
```
go build -o server_asr_llm_tts cmd/main.go
```


## Run server
Make sure to set the `DASHSCOPE_API_KEY` environment variable before running the server
```
export DASHSCOPE_API_KEY=XXXXX
./server_asr_llm_tts
```