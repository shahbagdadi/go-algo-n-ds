package main

import "fmt"

func sumEvenAfterQueries(nums []int, queries [][]int) []int {
	esum, ans := 0, []int{}
	for _, v := range nums {
		if v%2 == 0 {
			esum += v
		}
	}

	for _, q := range queries {
		v, idx := q[0], q[1]
		if nums[idx]%2 == 0 {
			esum -= nums[idx]
		}
		nums[idx] += v
		if nums[idx]%2 == 0 {
			esum += nums[idx]
		}
		ans = append(ans, esum)
	}
	return ans
}

func main() {
	nums := []int{1, 2, 3, 4}
	queries := [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}}
	ans := sumEvenAfterQueries(nums, queries)
	fmt.Println(ans)
	
}
