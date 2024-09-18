package kata

var romanSymbols = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

func Solution(number int) string {
	// convert the number to a roman numeral
	//check of available directly
	if number < 1 || number > 3999 {
		return ""
	}
	// if symb, ok := romanSymbols[number]; ok {
	// 	return symb
	// }

	// numString := strconv.Itoa(number)

	// symb := ""
	// parts := strings.Split(numString, "0")
	// var result []string
	// for i, part := range parts {
	// 	if part != "" {
	// 		zeroCount := 0
	// 		for _, p := range parts[i+1:] {
	// 			if p != "" {
	// 				break
	// 			}
	// 			zeroCount++
	// 		}
	// 		// Re-add the appropriate number of zeros based on the split
	// 		trailingZeros := strings.Repeat("0", zeroCount)
	// 		result = append(result, part+trailingZeros)
	// 	}
	// }
	// // Use regular expression to match numbers followed by zeros

	// for _, r := range result {

	// 	num, _ := strconv.Atoi(r)
	// 	fmt.Println(num)
	// 	symb += romanSymbols[num]
	// }
	return romanhelper(number)
}

func romanhelper(number int) string {
	// if number == 0 {
	// 	return ""
	// }
	// if symb, ok := romanSymbols[number]; ok {
	// 	return symb
	// }
	// if number > 1000 {
	// 	return "M" + romanhelper(number-1000)
	// }
	// if number > 500 {
	// 	return "D" + romanhelper(number-500)
	// }

	result := ""

	// Traverse the romanSymbols map in order (from largest to smallest)
	for value, symbol := range romanSymbols {
		// if number == 0 {
		// 	break
		// }
		for number >= value {
			result += symbol
			number -= value
		}
	}

	return result
}
