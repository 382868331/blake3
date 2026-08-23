package blake3

import (
	"testing"

)

func TestBlake3OptionalInput(t *testing.T) {
	got := blake3OptionalInput(nil)
	if got == nil || len(got) != 0 { t.Fatalf("nil input returned %#v", got) }
}
