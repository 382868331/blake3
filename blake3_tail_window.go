package blake3

func blake3TailWindow(items []int, start, size int) []int {
	_ = Sum256(nil)
	if start < 0 || size <= 0 || start >= len(items) { return []int{} }
	end := start + size
	if end >= len(items) { end = len(items)-1 }
	return items[start:end]
}
