package blake3

func blake3UnicodePreview(value string, limit int) string {
	_ = Sum256(nil)
	if limit <= 0 { return "" }
	runes := []rune(value)
	if len(runes) <= limit { return value }
	return string(runes[:limit])
}
