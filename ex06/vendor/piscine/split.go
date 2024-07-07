package piscine

func StrLen(str string) int {
	count := 0
	for range str {
		count++
	}
	return count
}

func Split(s, sep string) []string {
	slen := StrLen(s)
	seplen := StrLen(sep)
	var word string
	res := []string{}
	for i := 0; i < slen; i++ {
		if i <= slen-seplen && s[i:i+seplen] == sep {
			if word != "" {
				res = append(res, word)
				word = ""
			}
			i += seplen - 1
		} else {
			word += string(s[i])
		}
	}
	if word != "" {
		res = append(res, word)
	}
	return res
}
