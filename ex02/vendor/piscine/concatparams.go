package piscine

func ConcatParams(args []string) string {
	len := 0
	for range args {
		len++
	}
	res := ""
	for i, arg := range args {
		res += arg
		if i != len-1 {
			res += "\n"
		}
	}
	return res
}
