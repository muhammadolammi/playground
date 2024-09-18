package kata

import "math"

func FindNextSquare(sq int64) int64 {
	sR := math.Sqrt(float64(sq))
	// handle if its not a perfect square here
	if sR-float64(int(sR)) != 0 {
		return -1
	}
	return int64(math.Pow(sR+1, 2))
}
