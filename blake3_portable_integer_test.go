package blake3

import (
	"testing"

)

func TestBlake3PortableInteger(t *testing.T) {
	got,err:=blake3PortableInteger("4294967296")
	if err!=nil || got!=4294967296 { t.Fatalf("got=%d err=%v",got,err) }
}

func TestBlake3PortableIntegerRejectsInvalidText(t *testing.T) {
	if _,err:=blake3PortableInteger("4x"); err==nil { t.Fatal("expected parse error") }
}
