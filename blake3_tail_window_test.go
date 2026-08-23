package blake3

import (
	"testing"
	"reflect"
)

func TestBlake3TailWindow(t *testing.T) {
	got := blake3TailWindow([]int{1,2,3,4}, 2, 2)
	if !reflect.DeepEqual(got, []int{3,4}) { t.Fatalf("tail window = %v", got) }
}
