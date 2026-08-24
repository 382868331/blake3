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

func TestTaskBlake3006ContextSeparationPersistsAfterReset(t *testing.T) {
	left := blake3.NewDeriveKey("application/left")
	right := blake3.NewDeriveKey("application/right")
	left.Reset()
	right.Reset()
	_, _ = left.Write([]byte("shared material"))
	_, _ = right.Write([]byte("shared material"))
	if bytes.Equal(left.Sum(nil), right.Sum(nil)) {
		t.Fatal("reset discarded derivation context separation")
	}
}
