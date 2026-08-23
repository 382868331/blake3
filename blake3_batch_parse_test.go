package blake3

import (
	"testing"
	"reflect"
)

func TestBlake3BatchParse(t *testing.T) {
	got, err := blake3BatchParse([]string{"10","bad","30"})
	if err == nil || got != nil { t.Fatalf("got=%v err=%v", got, err) }
}
