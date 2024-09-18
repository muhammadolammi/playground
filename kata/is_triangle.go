package kata

/*
Given three sides with lengths
𝑎
a,
𝑏
b, and
𝑐
c, the conditions are:

𝑎+𝑏>𝑐

a+c>b
b+c>a
If all three conditions are satisfied, the three lengths can form a triangle.
*/

func IsTriangle(a, b, c int) bool {
	if a < 0 || b < 0 || c < 0 {
		return false
	}
	if a+b < c {
		return false
	}
	if a+c < b {
		return false
	}
	if b+c < a {
		return false
	}
	return true
}
