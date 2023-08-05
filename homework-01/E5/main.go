package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Read the input from STDIN
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, _ := strconv.Atoi(scanner.Text())

	// Fill out some base cases
	dp := []int{1, 2, 3, 5, 9}
	if input <= 5 {
		fmt.Println(dp[input-1])
		return
	}

	// Calculate dp
	n := 5
	for n != input {
		dp = append(dp, dp[n-1]+dp[n-2]+dp[n-5])
		n++
	}

	fmt.Println(dp[input-1])
}
