package dp

func jump(nums []int) int {
	/*
		1. Define Objective func
		minimum number of jump to nums[n-1]
		initially positioned at nums[0]
		and the number at nums[0] is the max jump possible to next jump
		2. Identify base case
		f(1) = f
		3. Recureence relation
		4. Order of computation
		5. Location of solution
	*/
	// greedy step can solve this
	jump := 0
	currentJumpEnd := 0
	farthest := 0
	for i := range len(nums) - 1 {
		farthest = max(farthest, i+nums[i])
		if i == currentJumpEnd {
			jump++
			currentJumpEnd = farthest
		}
		if currentJumpEnd > nums[len(nums)-1] {
			break
		}
	}

	return jump
}
