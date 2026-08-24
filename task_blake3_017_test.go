package blake3_test

import (
	"io"
	"testing"

	blake3 "github.com/zeebo/blake3"
)

func TestTaskBlake3017SeekCurrentPreservesPosition(t *testing.T) {
	h := blake3.New()
	d := h.Digest()
	prefix := make([]byte, 11)
	_, _ = d.Read(prefix)
	position, err := d.Seek(0, io.SeekCurrent)
	if err != nil {
		t.Fatal(err)
	}
	if position != int64(len(prefix)) {
		t.Fatalf("current position = %d, want %d", position, len(prefix))
	}
}

func TestTaskBlake3017SeekCurrentSupportsNegativeOffset(t *testing.T) {
	h := blake3.New()
	d := h.Digest()
	consumed := make([]byte, 70)
	_, _ = d.Read(consumed)
	position, err := d.Seek(-5, io.SeekCurrent)
	if err != nil {
		t.Fatal(err)
	}
	if position != 65 {
		t.Fatalf("current position = %d, want 65", position)
	}
}
