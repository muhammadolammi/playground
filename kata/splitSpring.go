package kata

func SplitString(str string) []string {
	words := []string{}
	end := 2
	start := 0
	for {
		// lets handle break
		if end > len(str) {
			break
		}
		word := str[start:end]
		if len(word) == 1 {
			word += "_"
		}
		words = append(words, word)

		end += 2
		// lets check if its the end
		if end > len(str) {
			// if its not an even lenght we subtract one from end. so it can just read one alphatbet in next run
			if len(str)%2 != 0 {
				end--
			}
		}
		start += 2
	}

	return words
}
