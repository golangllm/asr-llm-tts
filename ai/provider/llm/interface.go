package llm

import "context"

type LLMProvider interface {
	Stream(ctx context.Context, text <-chan string) (<-chan string, error)
}
