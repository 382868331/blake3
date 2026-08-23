package blake3

import (
	"testing"

)

func TestBlake3OptionalInput(t *testing.T) {
	got := blake3OptionalInput(nil)
	if got == nil || len(got) != 0 { t.Fatalf("nil input returned %#v", got) }
}

func TestBlake3OptionalInputCopyIsIndependent(t *testing.T) {
	in := []string{"a", "b"}; got := blake3OptionalInput(&in); got[0] = "changed"
	if in[0] != "a" { t.Fatalf("copy aliases input: %#v", in) }
}
