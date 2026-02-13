package tts

import "context"

// Chunk represents a piece of audio data.
type Chunk struct {
	Seq         int
	Audio       []byte // PCM16LE 16k mono by default
	Err         error
	Done        bool
	SampleRate  int
	AudioFormat string // e.g., "pcm16le"
}

// Streamer consumes text deltas and outputs audio chunks.
type Streamer interface {
	// Feed accepts a piece of text (can be partial incremental token).
	Feed(text string) error
	// Close indicates no more text will come.
	Close() error
	// Chunks returns the streaming audio chunks.
	Chunks() <-chan Chunk
}

// Factory creates a TTS stream.
type Factory interface {
	NewStream(ctx context.Context, voice string, sampleRate int) (Streamer, error)
}
