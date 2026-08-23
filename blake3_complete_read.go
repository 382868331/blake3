package blake3

import (
	"io"
)

func blake3CompleteRead(r io.Reader,size int) ([]byte,error) {
	buf:=make([]byte,size)
	_,err:=r.Read(buf)
	return buf,err
}
