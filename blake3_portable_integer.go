package blake3

import (
	"strconv"
)

func blake3PortableInteger(value string) (int64,error) {
	n,err:=strconv.ParseInt(value,10,32)
	return n,err
}
