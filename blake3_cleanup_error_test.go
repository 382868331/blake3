package blake3

import (
	"testing"
	"errors"
)

func TestBlake3CleanupError(t *testing.T) {
	want := errors.New("close failed")
	if err := blake3CleanupError(func() error{return nil}, func() error{return want}); !errors.Is(err,want) { t.Fatalf("err=%v",err) }
}

func TestBlake3CleanupErrorKeepsPrimaryError(t *testing.T) {
	want := errors.New("run failed")
	if err := blake3CleanupError(func() error{return want}, func() error{return errors.New("close")}); !errors.Is(err,want) { t.Fatalf("err=%v",err) }
}
