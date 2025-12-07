package util

import (
	"math/rand/v2"
)

// 抽奖。给定每个奖品被抽中的概率（无需要做归一化，但概率必须大于0），返回被抽中的奖品下标
func Lottery(probs []float64) int {
	if len(probs) == 0 {
		return -1
	}
	sum := 0.0
	acc := make([]float64, 0, len(probs)) //累积概率
	for _, prob := range probs {
		sum += prob
		acc = append(acc, sum)
	}

	// 获取(0,sum] 随机数
	r := rand.Float64() * sum
	index := BinarySearch4Section(acc, r)
	return index
}

func BinarySearch4Section(acc []float64, target float64) int {
	if len(acc) == 0 {
		return -1
	}

	left, right := 0, len(acc)-1
	ans := -1

	for left <= right {
		mid := (left + right) / 2
		if acc[mid] >= target {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}
