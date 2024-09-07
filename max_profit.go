package main

func maxProfit(prices []int) int {

	maxprofit := 0
	if len(prices) == 0 {
		return 0 // If it's empty, return 0 as there can be no profit
	}

	for i, sell := range prices[1:] {
		//lets check if if its good time to sel
		if sell > prices[i] {
			maxprofit += sell - prices[i]
		}

	}

	return maxprofit
}
