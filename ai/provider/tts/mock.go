package tts

import (
	"context"
	"errors"
	"math"
)

// MockStreamer turns any incoming text into short sine-wave beeps, just for end-to-end testing.
type MockStreamer struct {
	seq int
	ch  chan Chunk
}

func (s *MockStreamer) Feed(text string) error {
	// Generate ~120ms of a simple sine wave per call to simulate speech.
	// 16kHz mono PCM16LE
	const sampleRate = 16000
	const durationSamples = 1920 // 120ms
	const freq = 440.0
	buf := make([]byte, durationSamples*2)
	for i := 0; i < durationSamples; i++ {
		val := int16(3000 * math.Sin(2*math.Pi*freq*float64(i)/sampleRate))
		buf[2*i] = byte(val)
		buf[2*i+1] = byte(val >> 8)
	}
	s.seq++
	s.ch <- Chunk{Seq: s.seq, Audio: buf, SampleRate: sampleRate, AudioFormat: "pcm16le"}
	return nil
}

func (s *MockStreamer) Close() error {
	close(s.ch)
	return nil
}

func (s *MockStreamer) Chunks() <-chan Chunk { return s.ch }

// MockFactory creates Mock streamer.
type MockFactory struct{}

func (f MockFactory) NewStream(ctx context.Context, voice string, sampleRate int) (Streamer, error) {
	if ctx == nil {
		return nil, errors.New("nil context")
	}
	return &MockStreamer{ch: make(chan Chunk, 32)}, nil
}
