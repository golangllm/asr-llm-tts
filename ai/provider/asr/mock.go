package asr

import (
	"context"
	"errors"
)

// MockStreamer is a no-op ASR that never emits results unless explicitly fed special data.
// Useful for local development when no ASR service is configured.
type MockStreamer struct {
	ch chan Result
}

func (m *MockStreamer) PushAudio(pcm []byte) error { // ignore
	return nil
}

func (m *MockStreamer) Results() <-chan Result { return m.ch }

func (m *MockStreamer) Close() error {
	close(m.ch)
	return nil
}

// MockFactory creates mock streams.
type MockFactory struct{}

func (f MockFactory) NewStream(ctx context.Context, language string, sampleRate int) (Streamer, error) {
	if ctx == nil {
		return nil, errors.New("nil context")
	}
	return &MockStreamer{ch: make(chan Result)}, nil
}
