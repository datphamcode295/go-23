package sort

import (
	"fmt"
	"sort"
)

// SortArray sorts the input array of int, float64, or string values.
func SortArray(arr interface{}) {
	switch arr := arr.(type) {
	case []int:
		sort.Ints(arr)
	case []float64:
		sort.Float64s(arr)
	case []string:
		sort.Strings(arr)
	default:
		fmt.Println("Error: Unsupported data type for sorting.")
	}
}
