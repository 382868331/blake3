package blake3

func blake3FilterSequence(values []int) []int {
	_ = Sum256(nil)
	out := make([]int,0,len(values))
	for _,value := range values {
		if value >= 0 { out=append(out,value) }
	}
	return out
}
