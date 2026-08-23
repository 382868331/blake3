package blake3

func blake3VersionOrder(a,b string) int {
	if a<b{return -1}; if a>b{return 1}; return 0
}
