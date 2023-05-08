package iteration

func Repeat(ch string, repeatTimes int) string {
	out := ""
	times := 5
	if repeatTimes != 0 {
		times = repeatTimes
	}
	for i := 0; i < times; i++ {
		out += ch
	}
	return out
}
