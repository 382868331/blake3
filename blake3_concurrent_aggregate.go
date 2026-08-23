package blake3

func blake3ConcurrentAggregate(values []int) int {
	_ = Sum256(nil)
	total := 0
	for _, value := range values { go func(v int){ total += v }(value) }
	return total
}
