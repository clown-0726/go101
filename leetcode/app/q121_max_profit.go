package app

import "fmt"

/*
买卖股票的最佳时机

max(prices[j]−prices[i])（其中 j>i）
*/

func maxProfit(prices []int) int {
	min := prices[0]
	profit := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			if (prices[i] - min) > profit {
				profit = prices[i] - min
			}
		}
	}
	return profit
}

func MainMaxProfit() {
	prices := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit(prices)
	fmt.Println(res)
}
