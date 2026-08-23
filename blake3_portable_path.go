package blake3

import (
	"path"
)

func blake3PortablePath(base,name string) string {
	if base=="" { return path.Clean("/"+name) }
	return path.Join(base,name)
}
