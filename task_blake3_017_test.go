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
