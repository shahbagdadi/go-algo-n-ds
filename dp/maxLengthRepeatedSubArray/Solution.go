package main

import "fmt"

func findLength(nums1 []int, nums2 []int) int {
    ans, L1, L2 := 0, len(nums1) , len(nums2)
	prev := make([]int,L2+1)
	for i := 0; i < L1; i++ {
		curr := make([]int,L2+1)
		for j := 0; j < L2; j++ {
			if nums1[i] == nums2[j] {
				curr[j+1] = prev[j] + 1 
				ans = Max(ans,curr[j+1])
			}
		}
		prev = curr
	}
	return ans
}

func Max(x, y int) int {
	if x > y {
	  return x
	}
	return y
   }


func main() {
	ip1 := []int{1, 2, 3, 2, 1}
	ip2 := []int{3,2,1,4,7}
	ans := findLength(ip1,ip2)
	fmt.Println(ans)
}