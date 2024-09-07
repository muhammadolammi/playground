package dp

/*
problem: return sum of all n from i =1 to n

f(n) = f(n-1) + i
f(0) = 0 //

	if we do f(0-1) it will be equal to -1 and there is no -1 index, so we start our loop from 1
	if n == 0 , we can just return zero early.
*/
func Sum(n int) int {
	// the size here is plus 1 because we want our index to be in 1 index, the first value here dp[0] will be auto set to zero.

	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + i

	}
	return dp[n]

	// the up one is learning dp, i feel for long number the list created will be two much, so that space is a lot
	// let use the aritmetic sum formular s= n/2 x (a+l)
	// a =1
	// l=n

	// sum := float64(n) / 2 * float64(1+n)
	// return int(sum)

	// these two line of code is sufficient
	// running a loop for number 1000000000000 is too much, it will run out of memory
}
