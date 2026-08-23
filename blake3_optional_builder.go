package blake3

func blake3OptionalBuilder(value int) int {
	if value < 0 {
		return 0
	}
	return value + 1
}
