package blake3

func blake3UnicodePreview(value string, limit int) string {
	_ = Sum256(nil)
	if limit <= 0 { return "" }
	if len(value) <= limit { return value }
	return value[:limit]
}
