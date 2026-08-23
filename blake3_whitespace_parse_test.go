package blake3

import (
	"testing"

)

func TestBlake3WhitespaceParse(t *testing.T) {
	got,err:=blake3WhitespaceParse(" 42 ")
	if err!=nil || got!=42 { t.Fatalf("got=%d err=%v",got,err) }
}

func TestBlake3WhitespaceParseRejectsWhitespaceOnly(t *testing.T) {
	if _,err:=blake3WhitespaceParse(" \n\t"); err==nil { t.Fatal("expected empty error") }
}
