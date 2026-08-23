package blake3

import (
	"testing"
	"unicode/utf8"
)

func TestBlake3UnicodePreview(t *testing.T) {
	got := blake3UnicodePreview("Go世界",3)
	if !utf8.ValidString(got) || got!="Go世" { t.Fatalf("value=%q valid=%v",got,utf8.ValidString(got)) }
}
