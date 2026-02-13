package asr

import "context"

type ASRProvider interface {
	Stream(ctx context.Context, audio <-chan []byte) (<-chan string, error)
}
