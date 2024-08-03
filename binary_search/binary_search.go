// Package binary_search provide two function to perform binary search. Iterative and recursive way

package binary_search

// BinarySearchIterative recieves an array of int and a key of int to search for then return an int of the index where the key was found
// This function works iteratively
// The Time of this func is O(logn) to test this uncomment the comment and comment your return statemnt in for
func BinarySearchIterative(a []int, key int) int {
	l := 0
	h := len(a) - 1
	// index := -1
	// count := 0

	for l <= h {

		middle := (l + h) / 2

		if a[middle] == key {
			// index = middle
			// break
			return middle
		}
		// count++
		if key > a[middle] {
			l = middle + 1
		}
		if key < a[middle] {
			h = middle - 1
		}

	}
	// fmt.Printf("ran for %d times\n", count)
	return -1

}

// BinarySearch recieves an array of int and a key of int to search for then return an int of the index where the key was found
// This function works recursively
func BinarySearch(a []int, key, l, h int) int {
	if l == h {
		if a[l] == key {
			return l
		}
		return -1

	}
	mid := (l + h) / 2
	if a[mid] == key {
		return mid
	}
	if key < a[mid] {
		return BinarySearch(a, key, l, mid-1)
	}
	return BinarySearch(a, key, mid+1, h)

}
