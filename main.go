package main

import (
	"fmt"
	"sort"
)

// 算法
// 1.对2个数组进行从小到大排序，这里用go的“sort”包进行排序
// 2.2层循环，外层是数组1，内层是数组2，
// 3.跳出外层循环的条件是遍历完数组，或者外层当前值加上内层最小值大于预算
// 4.内层索引移动规则为，
//
//	4.1如果开始的时候是在最左端，那么向右移动，直到加上外层的大于预算或者到最右
//	4.2如果开始的时候在最左端，那么向左移动，直到加上外层的小于等于预算或者到最左
//	4.3如果开始的时候在中间，那么向左移动，直到加上外层小于等于预算或者到最左
func OptionNumber(DevPrices []int, InsurePrices []int, sum int) int {
	// 时间复杂度为 log(m) + log(n)
	sort.Slice(DevPrices, func(i, j int) bool {
		return DevPrices[i] <= DevPrices[j]
	})
	sort.Slice(InsurePrices, func(i, j int) bool {
		return InsurePrices[i] <= InsurePrices[j]
	})
	// 取边界值
	if DevPrices[len(DevPrices)-1]+InsurePrices[len(InsurePrices)-1] <= sum {
		return len(DevPrices) * len(InsurePrices)
	}

	// 时间复杂度平均为 m*n/4
	ret := 0
	indexInsure := 0
	for _, devPrice := range DevPrices {
		if indexInsure == 0 {
			if devPrice+InsurePrices[indexInsure] > sum {
				break
			}
			for j := indexInsure; j < len(InsurePrices); j++ {
				if devPrice+InsurePrices[j] <= sum {
					ret += 1
					indexInsure += 1
				} else {
					break
				}
			}
		} else {
			if devPrice+InsurePrices[indexInsure] <= sum {
				ret += len(InsurePrices)
				break
			} else {
				for j := len(InsurePrices); j >= 0; j-- {
					indexInsure = j
					if devPrice+InsurePrices[j] <= sum {
						ret += j + 1
						break
					}
				}
			}
		}
	}
	return ret

}

func main() {
	fmt.Println(OptionNumber([]int{10, 20, 5}, []int{5, 5, 2}, 15))
}
