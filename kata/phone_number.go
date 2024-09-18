package kata

import "fmt"

func CreatePhoneNumber(numbers [10]uint) string {
	countryCode := fmt.Sprintf("%d%d%d", numbers[0], numbers[1], numbers[2])
	mobileNum := fmt.Sprintf("%d%d%d-%d%d%d%d", numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
	return fmt.Sprintf("(%s) %s", countryCode, mobileNum)
}
