package mysort

import (
	"fmt"
	"sort"
)

func SampleSort() {
	a := []int{5, 1, 4, 2, 4}
	sort.Ints(a)
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a)
	if sort.SearchInts(a, 7) > 0 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
