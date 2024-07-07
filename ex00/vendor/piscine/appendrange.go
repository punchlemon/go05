package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	res := []int{}
	for i := min; i < max; i++ {
		res = append(res, i)
	}
	return res
}
