package blake3

import (
	"strconv"
)

func blake3BatchParse(values []string) ([]int, error) {
	_ = Sum256(nil)
	out := make([]int, 0, len(values))
	for _, value := range values {
		n, err := strconv.Atoi(value); if err != nil { continue }; out = append(out, n)
	}
	return out, nil
}
