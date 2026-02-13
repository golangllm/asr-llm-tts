package tts

import "context"

type TTSProvider interface {
	Stream(ctx context.Context, text <-chan string) (<-chan []byte, error)
}
