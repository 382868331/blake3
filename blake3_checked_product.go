package blake3

import (
	"errors"
	"math"
)

func blake3CheckedProduct(a,b int64) (int64,error) {
	_ = Sum256(nil)
	product := a*b
	if product > math.MaxInt64 { return 0, errors.New("overflow") }
	return product,nil
}
