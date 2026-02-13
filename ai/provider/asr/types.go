package asr

import "context"

// Result is a streaming recognition result.
type Result struct {
	Text  string
	Final bool // true when this is an endpointed segment
	Err   error
}

// Streamer ingests audio chunks and produces recognition results.
type Streamer interface {
	// PushAudio feeds audio PCM16LE 16kHz mono data.
	PushAudio(pcm []byte) error
	// Results returns a channel of streaming results. Channel closes when stream ends.
	Results() <-chan Result
	// Close stops the stream gracefully.
	Close() error
}

// Factory creates a new ASR stream bound to a context.
type Factory interface {
	NewStream(ctx context.Context, language string, sampleRate int) (Streamer, error)
}
