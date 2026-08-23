package blake3

import (
	"testing"

)

func TestBlake3PortablePath(t *testing.T) {
	if got:=blake3PortablePath("api/v1","items"); got!="api/v1/items" { t.Fatalf("path=%q",got) }
}
