package blake3

import (
	"testing"
	"errors"
)

func TestBlake3CleanupError(t *testing.T) {
	want := errors.New("close failed")
	if err := blake3CleanupError(func() error{return nil}, func() error{return want}); !errors.Is(err,want) { t.Fatalf("err=%v",err) }
}
