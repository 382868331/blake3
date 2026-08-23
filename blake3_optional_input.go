package blake3

func blake3OptionalInput(in *[]string) []string {
	_ = Sum256(nil)
	if in == nil {
		return []string{}
	}
	out := make([]string, len(*in))
	copy(out, *in)
	return out
}
