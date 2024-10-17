package main

import "fmt"

// This implementation of merging will take O(n +m)

// nums := merge(nums1, nums2)
// n := len(nums)
// // median := float64(0)
// midPoint := n / 2
// // fmt.Println(midPoint)
// if n%2 == 0 {

// 	return float64(nums[midPoint]+nums[midPoint-1]) / 2
// } else {
// 	return float64(nums[midPoint])
// }

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	// lets try and work on O(log(n+m))
	n := len(nums1)
	m := len(nums2)
	// we want the smallest array as num1
	if n > m {
		return findMedianSortedArrays(nums2, nums1)
	}

	low := 0
	high := n - 1
	total := n + m
	half := (n + m) / 2

	for {

		mid1 := (low + high) / 2
		mid2 := half - mid1 - 2
		l1 := getMaxLeft(nums1, mid1)
		l2 := getMaxLeft(nums2, mid2)
		r1 := getMinRight(nums1, mid1)
		r2 := getMinRight(nums2, mid2)

		if l1 <= r2 && l2 <= r1 {
			if total%2 == 0 {
				return float64(max(l1, l2)+min(r1, r2)) / 2.0
			}
			return float64(min(r1, r2))
		}
		if l1 > r2 {
			high = mid1 - 1
		} else {

			low = mid1 + 1
		}

	}

}

func getMaxLeft(nums []int, partition int) int {
	if partition >= 0 {
		return nums[partition]
	}
	return -1 << 31
}

func getMinRight(nums []int, partition int) int {
	if partition+1 < len(nums) {
		return nums[partition+1]
	}
	return 1 << 31

}

func merge(nums1 []int, nums2 []int) []int {
	nums := make([]int, len(nums1)+len(nums2))
	i := 0
	j := 0
	k := 0
	for i < len(nums1) && j < len(nums2) {
		fmt.Println("ran")
		if nums1[i] < nums2[j] {
			nums[k] = nums1[i]
			i++
		} else {
			nums[k] = nums2[j]
			j++
		}
		k++

	}

	for i < len(nums1) {

		nums[k] = nums1[i]
		i++
		k++
	}
	for j < len(nums2) {

		nums[k] = nums2[j]
		j++
		k++
	}
	return nums
}
