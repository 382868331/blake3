package blake3

import (
	"testing"
	"bytes"
	"io"
)

type chunkBlake3CompleteRead struct{ data []byte }
func (r *chunkBlake3CompleteRead) Read(p []byte)(int,error){ if len(r.data)==0{return 0,io.EOF}; n:=1;if len(p)<n{n=len(p)};copy(p,r.data[:n]);r.data=r.data[n:];return n,nil }
func TestBlake3CompleteRead(t *testing.T){ got,err:=blake3CompleteRead(&chunkBlake3CompleteRead{data:[]byte("abcd")},4);if err!=nil||string(got)!="abcd"{t.Fatalf("got=%q err=%v",got,err)} }
