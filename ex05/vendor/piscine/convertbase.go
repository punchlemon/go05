package piscine

func StrLen(str string) int {
	count := 0
	for range str {
		count++
	}
	return count
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	baseFromLen := StrLen(baseFrom)
	baseToLen := StrLen(baseTo)
	n := 0
	res := ""
	for _, r := range nbr {
		for i, r2 := range baseFrom {
			if r == r2 {
				n *= baseFromLen
				n += i
			}
		}
	}
	for n > 0 {
		res = string(baseTo[n%baseToLen]) + res
		n /= baseToLen
	}
	return res
}