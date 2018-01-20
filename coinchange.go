package main

import "fmt"

var mem [51][251]int
var present [51][251]bool

func ways(i, target int, coins []int) int {

	if target == 0 {
		return 1
	}

	if target < 0 {
		return 0
	}

	if i < 0 {
		return 0
	}

	if present[i][target] == true {
		return mem[i][target]
	} else {
		total_ways := ways(i, target - coins[i], coins) + ways(i - 1, target, coins)
		mem[i][target] = total_ways
		present[i][target] = true
		return total_ways
	}
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var target, coin_count int
	fmt.Scan(&target)
	fmt.Scan(&coin_count)

	var coins = make([]int, coin_count)
	for i := 0; i < coin_count; i++ {
		fmt.Scan(&coins[i])
	}

	fmt.Println(ways(len(coins) - 1, target, coins[:]))
}