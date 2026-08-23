package blake3

import (
	"testing"

)

func TestBlake3VersionOrder(t *testing.T) {
	if got:=blake3VersionOrder("1.10","1.9"); got<=0 { t.Fatalf("comparison=%d",got) }
}

func TestBlake3VersionOrderTreatsMissingSegmentsAsZero(t *testing.T) {
	if got:=blake3VersionOrder("2.0","2"); got!=0 { t.Fatalf("comparison=%d",got) }
}
