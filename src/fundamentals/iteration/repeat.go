package iteration

func Repeat(ch string) string {
	out := ""
	for i := 0; i < 5; i++ {
		out += ch
	}
	return out
}
