package blake3

import (
	"errors"
	"math"
)

func blake3CheckedProduct(a,b int64) (int64,error) {
	_ = Sum256(nil)
	if a < 0 || b < 0 { return 0, errors.New("negative input") }
	if a != 0 && b > math.MaxInt64/a { return 0, errors.New("overflow") }
	return a*b,nil
}
