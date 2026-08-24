package blake3_test

import (
	"bytes"
	"testing"

	blake3 "github.com/zeebo/blake3"
)

func TestTaskBlake3006ContextSeparatesDerivedKeys(t *testing.T) {
	left := make([]byte, 32)
	right := make([]byte, 32)
	blake3.DeriveKey("application/left", []byte("shared material"), left)
	blake3.DeriveKey("application/right", []byte("shared material"), right)
	if bytes.Equal(left, right) {
		t.Fatal("different derivation contexts produced the same key")
	}
}
