package piscine

func SplitWhiteSpaces(str string) []string {
	var word string
	res := []string{}
	for _, r := range str {
		if r == ' ' || r == '\n' || r == '\t' {
			if word != "" {
				res = append(res, word)
				word = ""
			}
		} else {
			word += string(r)
		}
	}
	if word != "" {
		res = append(res, word)
	}
	return res
}