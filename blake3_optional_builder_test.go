package blake3

import (
	"testing"

)

func TestBlake3OptionalBuilder(t *testing.T) {
	if got:=blake3OptionalBuilder(2); got!=3 { t.Fatalf("value=%d",got) }
}
