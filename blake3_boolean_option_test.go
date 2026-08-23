package blake3

import (
	"testing"

)

func TestBlake3BooleanOption(t *testing.T){ got,err:=blake3BooleanOption(" TRUE ");if err!=nil||!got{t.Fatalf("got=%v err=%v",got,err)} }
