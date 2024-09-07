package quicksort

import "fmt"

func swap(Array []int, index1, index2 int) {
	temp := Array[index2]
	Array[index2] = Array[index1]
	Array[index1] = temp
	fmt.Println("ran")
}
func partition(Array []int, l, h int) int {
	// pivot := Array[h]
	// i := l
	// for j := l; j < h; j++ {
	// 	if Array[j] < pivot {

	// 		swap(Array, i, j)
	// 		i++
	// 	}
	// }
	// swap(Array, i, h)
	// return i
	pivot := Array[l]
	i := l
	j := h
	for i < j {
		// Get the number greater than the pivot
		for Array[j] > pivot {
			j--
		}
		// Get the number less than the pivot

		for Array[i] <= pivot {
			i++
		}
		// Make sure the i < j , making sure the pivot hasnt reach it sorted position
		if i < j {
			swap(Array, i, j)
		}
	}
	// If it reach here it means it has reach the pivot sorted position

	swap(Array, l, j)
	//    Return the pivot sorted position so we can call recursive call on the 2 parts
	return j

	// THIS ALGORITHM IS SHIT, O(n^2 ) very slow compared to mergesort
}
func QuickSort(Array []int, l, h int) {
	if l < h {
		j := partition(Array, l, h)
		QuickSort(Array, l, j)
		QuickSort(Array, j+1, h)
	}
}
