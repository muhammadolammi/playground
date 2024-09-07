package main

import "fmt"

func removeDuplicates(nums []int) int {
	numsRecord := map[int]bool{}
	count := 0

	for i, num := range nums {
		if !numsRecord[num] {
			//
			numsRecord[num] = true
			nums[count] = nums[i]
			count++
		}

	}
	fmt.Println(nums)
	return count
}

func removeDuplicates2(nums []int) int {
	numsRecord := map[int][]int{}
	count := 0

	for i, num := range nums {
		if len(numsRecord[num]) != 2 {
			// 0, 0, 1, 1, 1, 2, 2, 3, 3, 4
			numsRecord[num] = append(numsRecord[num], num)
			nums[count] = nums[i]
			count++
		}

	}
	return count
}
