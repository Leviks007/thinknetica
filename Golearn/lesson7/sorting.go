package lesson7

import "sort"

func sortINT(x []int) []int {
	sort.Ints(x)
	return x
}

func sortString(x []string) []string {
	sort.Strings(x)
	return x
}

func sortfloat64(x []float64) []float64 {
	sort.Float64s(x)
	return x
}
