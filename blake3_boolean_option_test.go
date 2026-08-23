package blake3

import (
	"testing"

)

func TestBlake3BooleanOption(t *testing.T){ got,err:=blake3BooleanOption(" TRUE ");if err!=nil||!got{t.Fatalf("got=%v err=%v",got,err)} }

func TestBlake3BooleanOptionRejectsUnknownValue(t *testing.T){ if _,err:=blake3BooleanOption("enabled");err==nil{t.Fatal("expected validation error")} }
