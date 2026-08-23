package blake3

import (
	"testing"
	"unicode/utf8"
)

func TestBlake3UnicodePreview(t *testing.T) {
	got := blake3UnicodePreview("Go世界",3)
	if !utf8.ValidString(got) || got!="Go世" { t.Fatalf("value=%q valid=%v",got,utf8.ValidString(got)) }
}

func TestBlake3UnicodePreviewKeepsShortText(t *testing.T) {
	if got:=blake3UnicodePreview("世界",5); got!="世界" { t.Fatalf("value=%q",got) }
}
