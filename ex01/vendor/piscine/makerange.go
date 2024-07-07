package piscine

func MakeRange(min, max int) []int {
	len := max - min
	if len <= 0 {
		return nil
	}
	res := make([]int, len)
	for i := 0; i < len; i++ {
		res[i] = i + min
	}
	return res
}
