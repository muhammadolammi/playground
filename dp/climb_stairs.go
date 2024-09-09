package dp

import (
	"fmt"
	"math"
)

func ClimbStairs(n int) int {
	/*
		1. Define Objective func
		The objective is to return number of distinct ways we can reach the nth stair. we are allowed to move a step or two per stairs

		2. Identify base case
		f(0) = 1
		f(1)=1
		/ these two f(0) and f(1) are bases
		f(2) = 2
		f(3)= 3
		f(4) = 5
		f(n) = f(n-1)+ f(n-2)
		3. Recureence relation
		f(n) = f(n-1)+ f(n-2)
		4. Order of computation
		0 -> n
		5. Location of solution
		f(n)

	*/

	// dp := make([]int, 2)
	// dp[0] = 1
	// dp[1] = 1
	// for i := 2; i <= n; i++ {
	// 	dp[i] = dp[i-1] + dp[i-2]
	// }
	// return dp[n]

	// with space optimzation
	a := 1
	b := 1
	c := 0
	for i := 2; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c

}

func ClimbStairs3Steps(n int) int {
	/*
		1. Define Objective func
		The objective is to return number of distinct ways we can reach the nth stair.
		we are allowed to move a step or two or three per stairs

		2. Identify base case
		f(0) = 1
		f(1)=1
		/ these two f(0) and f(1) are bases
		f(2) = 2
		f(3)= 4
		f(4) = 7
		f(n) = f(n-1)+ f(n-2) + f(n-3)
		3. Recureence relation
		f(n) = f(n-1)+ f(n-2) + f(n-3)
		4. Order of computation
		0 -> n
		5. Location of solution
		f(n)

	*/

	// dp := make([]int, 2)
	// dp[0] = 1
	// dp[1] = 1
	// for i := 2; i <= n; i++ {
	// 	dp[i] = dp[i-1] + dp[i-2]
	// }
	// return dp[n]

	// with space optimzation
	a := 1
	b := 1
	c := 2
	d := 0
	for i := 3; i <= n; i++ {
		d = a + b + c
		a = b
		b = c
		c = d
	}
	return d

}

func ClimbStairsKSteps(n, k int) int {
	/*
		1. Define Objective func
		The objective is to return number of distinct ways we can reach the nth stair.
		we are allowed to move a step or two or three per stairs

		2. Identify base case
		f(0) = 1
		f(1)=1
		/ these two f(0) and f(1) are bases
		f(2) = 2
		f(3)= 4
		f(4) = 7
		f(n) = f(n-1)+ f(n-2) + f(n-3) .... f(n-k)
		3. Recureence relation
		f(n) = f(n-1)+ f(n-2) + f(n-3)
		4. Order of computation
		0 -> n
		5. Location of solution
		f(n)

	*/

	// dp := make([]int, 2)
	// dp[0] = 1
	// dp[1] = 1
	// for i := 2; i <= n; i++ {
	// 	dp[i] = dp[i-1] + dp[i-2]
	// }
	// return dp[n]

	// with space optimzation
	dp := make([]int, k)
	dp[0] = 1
	// dp[1] = 1
	// [0, 0, 0]
	// [1, 1, 2]
	// [4, 7, 13]
	// [24, 44, 81]
	for i := 1; i <= n; i++ {
		for j := 1; j < k; j++ {
			if i-j < 0 {
				continue
			}
			dp[i%k] += dp[(i-j)%k]
		}
	}
	return dp[n%k]

}

func PaidStairCases(n, k int, p []int) int {
	/*
				1. Define Objective func
				Return the minimum price needed to react stair n
				/ this is for 2 steps and same idea will work for  k steps
				f(i) = p(i) + min(p(i-1), p(i-2), p(i-3).... p(i-k))
				2. Identify base case
				f(0)=0
		        f(1)= p[1]
				3. Recureence relation
				f(i) = p(i) + min(p(i-1), p(i-2), p(i-3).... p(i-k))
				4. Order of computation
				bottom up
				5. Location of solution
				f(n)
	*/

	dp := make([]int, k)
	dp[0] = 0

	for i := 1; i <= n; i++ {
		m := math.MaxInt

		// lets range trough k
		for j := 1; j <= k; j++ {
			// lets check when we dont need to run up to k
			if i-j < 0 {
				break
			}
			m = min(m, dp[(i-j)%k])
		}
		dp[i%k] = p[i] + m

	}

	return dp[n%k]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func PaidStairCasesPath(n int, p []int) []int {
	/*
				1. Define Objective func
				Return the minimum price needed to react stair n
				/ this is for 2 steps and same idea will work for  k steps
				f(i) = p(i) + min(p(i-1), p(i-2), p(i-3).... p(i-k))
				2. Identify base case
				f(0)=0
		        f(1)= p[1]
				3. Recureence relation
				f(i) = p(i) + min(p(i-1), p(i-2), p(i-3).... p(i-k))
				4. Order of computation
				bottom up
				5. Location of solution
				f(n)

				/ we are to return a path here
	*/

	dp := make([]int, n+1)
	from := make([]int, n+1)
	dp[0] = 0
	dp[1] = p[1]

	for i := 2; i <= n; i++ {

		dp[i] = p[i] + min(dp[i-1], dp[i-2])
		if dp[i-1] < dp[i-2] {
			from[i] = i - 1
		} else {
			from[i] = i - 2
		}

	}

	fmt.Println(dp)
	fmt.Println(from)
	path := []int{}
	for cur := n; cur > 0; cur = from[cur] {
		path = append(path, cur)
	}
	fmt.Println(path)
	return path
}
