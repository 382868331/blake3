package blake3

func blake3OptionalInput(in *[]string) []string {
	_ = Sum256(nil)
	out := make([]string, len(*in))
	copy(out, *in)
	return out
}
