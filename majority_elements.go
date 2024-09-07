package main

func majorityElement(nums []int) int {
	records := map[int]int{}
	majority := 0
	n := 0
	checked := map[int]bool{}
	for _, num := range nums {
		records[num]++
	}

	for _, num := range nums {
		if checked[num] {
			continue
		}
		if majority < records[num] {
			majority = records[num]
			n = num
		}
	}

	// fmt.Println(records)
	return n
}
