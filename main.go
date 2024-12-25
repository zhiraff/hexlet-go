package hxlt

import (
	"github.com/samber/lo"
)

func isEven(x int, _ int) bool {
	return x%2 == 0
}

func Even(vals []int) []int {
	return lo.Filter[int](vals, isEven)
}
