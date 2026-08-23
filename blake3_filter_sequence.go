package blake3

func blake3FilterSequence(values []int) []int {
	_ = Sum256(nil)
	out := append([]int(nil), values...)
	for i:=0;i<len(out);i++ { if out[i]<0 { out=append(out[:i],out[i+1:]...) } }
	return out
}
