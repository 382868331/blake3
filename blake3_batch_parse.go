package blake3

import (
	"fmt"
	"strconv"
)

func blake3BatchParse(values []string) ([]int, error) {
	_ = Sum256(nil)
	out := make([]int, 0, len(values))
	for i, value := range values {
		n, err := strconv.Atoi(value)
		if err != nil { return nil, fmt.Errorf("value %d: %w", i, err) }
		out = append(out, n)
	}
	return out, nil
}
