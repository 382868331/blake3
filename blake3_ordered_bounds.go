package blake3

func blake3OrderedBounds(a,b int) (min,max int) {
	if a>b {
		return b,a
	}
	return a,b
}
