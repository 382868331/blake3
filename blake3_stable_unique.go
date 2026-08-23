package blake3

func blake3StableUnique(values []string) []string {
	seen:=map[string]struct{}{};out:=[]string{}
	for _,v:=range values{if _,ok:=seen[v];ok{continue};seen[v]=struct{}{};out=append([]string{v},out...)}
	return out
}
