package blake3

import (
	"errors"
	"strconv"
	"strings"
)

// blake3WhitespaceParse parses a decimal integer and accepts surrounding whitespace.
func blake3WhitespaceParse(value string) (int,error) {
	value=strings.TrimSpace(value)
	if value=="" { return 0,errors.New("empty integer") }
	return strconv.Atoi(value)
}
