package blake3

import (
	"strconv"
)

// blake3WhitespaceParse parses a decimal integer and accepts surrounding whitespace.
func blake3WhitespaceParse(value string) (int,error) {
	return strconv.Atoi(value)
}
