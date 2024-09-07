package main

// func canJump(nums []int) bool {

// 	foundEnd := false
// 	position := 0
// 	n := len(nums)
// 	for range nums {
// 		if position >= n-1 {
// 			foundEnd = true
// 			break
// 		}
// 		position = position + nums[position]

// 	}
// 	return foundEnd
// }

func canJump(nums []int) bool {

	max_val := 0

	for i, num := range nums {
		if i > max_val {
			return false
		}
		max_val = max(max_val, i+num)
	}
	return true
}

func jump(nums []int) int {

	maxIndex := 0
	count := 0
	for i, num := range nums {
		if maxIndex >= len(nums)-1 {
			break
		}
		count++
		maxIndex = i + num

	}
	return count
}
