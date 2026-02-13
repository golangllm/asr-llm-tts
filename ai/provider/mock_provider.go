package provider

import (
	"context"
	"time"
)

type MockASR struct{}

func (m *MockASR) Stream(ctx context.Context, audio <-chan []byte) (<-chan string, error) {
	out := make(chan string)
	go func() {
		defer close(out)
		for range audio {
			time.Sleep(200 * time.Millisecond)
			out <- "你好，这是 ASR 文本"
		}
	}()
	return out, nil
}

type MockLLM struct{}

func (m *MockLLM) Stream(ctx context.Context, text <-chan string) (<-chan string, error) {
	out := make(chan string)
	go func() {
		defer close(out)
		for t := range text {
			for _, r := range []rune("【LLM 回复】" + t) {
				time.Sleep(100 * time.Millisecond)
				out <- string(r)
			}
		}
	}()
	return out, nil
}

type MockTTS struct{}

func (m *MockTTS) Stream(ctx context.Context, text <-chan string) (<-chan []byte, error) {
	out := make(chan []byte)
	go func() {
		defer close(out)
		for t := range text {
			time.Sleep(150 * time.Millisecond)
			out <- []byte("AUDIO(" + t + ")")
		}
	}()
	return out, nil
}
