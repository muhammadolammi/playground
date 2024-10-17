package main

import (
	"math"
)

// lets use sliding window
func lengthOfLongestSubstring(s string) int {
	left := 0
	window := map[byte]bool{}
	maxLen := 0
	for right := 0; right < len(s); right++ {

		for window[s[right]] {
			delete(window, s[left])
			left++
		}
		window[s[right]] = true

		maxLen = max(maxLen, right-left+1)
	}
	return maxLen
}

func findMaxSubArray(arr []int, k int) int {
	if len(arr) < k {
		return 0
	}
	maxSum := 0
	currentSum := 0
	for i := 0; i < k; i++ {
		currentSum += arr[i]
	}
	if len(arr) == k {
		return currentSum
	}

	// Now lets slide if the array > k
	for i := k; i < len(arr); i++ {
		currentSum -= arr[i-k]
		currentSum += arr[i]
		maxSum = max(currentSum, maxSum)
	}

	return maxSum

}

// this func find the smallest sub array with sum>= x
func smallestSubArrayGivenSum(arr []int, x int) int {
	smallestSubArray := math.MaxInt
	currentSum := 0
	left := 0
	for right := 0; right < len(arr); right++ {
		if smallestSubArray == 1 {
			break
		}
		currentSum += arr[right]

		for currentSum >= x {
			smallestSubArray = min(smallestSubArray, right-left+1)
			currentSum -= arr[left]
			left++
		}

	}

	return smallestSubArray
}
