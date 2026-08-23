package blake3

import (
	"strconv"
)

func blake3PortableInteger(value string) (int64,error) {
	n,err:=strconv.ParseInt(value,10,64)
	if err!=nil { return 0,err }
	return n,nil
}
