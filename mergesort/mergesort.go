package mergesort

func Merge(A, B []int) []int {
	C := make([]int, len(A)+len(B))
	i := 0
	j := 0
	k := 0
	for i < len(A) && j < len(B) {
		if A[i] < B[j] {
			C[k] = A[i]
			i++
			k++
			continue
		}

		C[k] = B[j]
		j++
		k++

	}

	for i < len(A) {
		C[k] = A[i]
		i++
		k++
	}
	for j < len(B) {
		C[k] = B[j]
		j++
		k++
	}

	return C
}

func MergeM(A, B, C, D []int) []int {
	l1 := Merge(A, B)
	l2 := Merge(C, D)
	return Merge(l1, l2)
}

func MergeSort(A []int) []int {
	if len(A) < 2 {
		return A
	}
	first := MergeSort(A[:len(A)/2])

	second := MergeSort(A[len(A)/2:])
	return Merge(first, second)

}
