package blake3

import (
	"testing"
	"reflect"
)

func TestBlake3StableUnique(t *testing.T){ want:=[]string{"b","a","c"};got:=blake3StableUnique([]string{"b","a","b","c"});if !reflect.DeepEqual(got,want){t.Fatalf("order=%v",got)} }
