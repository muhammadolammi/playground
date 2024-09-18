package kata

func FindUniq(arr []float32) float32 {
	// Do the magic
	// lets see if we can see in the first 3
	// if arr[0] != arr[1] || arr[1] != arr[2] {

	// 	if arr[1] == arr[2] {
	// 		fmt.Println("ran")
	// 		fmt.Println(arr[0])
	// 		return arr[0]
	// 	}
	// 	fmt.Println(arr[1])
	// 	return arr[1]

	// }
	// // if it reach here , the first 3 are the same
	// nonUnique := arr[2]
	// unique := float32(0)
	// for _, n := range arr[3:] {
	// 	if n != nonUnique {
	// 		unique = n
	// 	}
	// }

	// return unique

	un, found := helper(arr[:3])
	if found {
		return un
	}
	un, _ = helper(arr[3:])
	return un

}

func helper(array []float32) (float32, bool) {
	n1 := array[0]
	n2 := array[1]
	n3 := array[2]
	if n1 != n2 || n1 != n3 || n2 != n3 {
		if n1 != n2 && n1 != n3 {
			return n1, true
		}
		if n2 != n1 && n2 != n3 {
			return n2, true
		}
		return n3, true
	}

	if len(array) <= 3 {
		return float32(0), false
	}
	nonUnique := float32(0)
	if array[0] == array[1] {
		nonUnique = array[0]
	}
	unique := float32(0)
	status := false

	for _, n := range array {
		if n != nonUnique {
			status = true
			unique = n
		}
	}

	return unique, status

}
