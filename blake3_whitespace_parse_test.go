package blake3

import (
	"testing"

)

func TestBlake3WhitespaceParse(t *testing.T) {
	got,err:=blake3WhitespaceParse(" 42 ")
	if err!=nil || got!=42 { t.Fatalf("got=%d err=%v",got,err) }
}
