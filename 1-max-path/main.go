package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var filenames = []string{
	"easy.json",
	"hard.json",
}

type TestCase [][]int

func main() {
	for _, f := range filenames {
		bytes, err := os.ReadFile(f)
		if err != nil {
			panic(err)
		}

		var tc TestCase
		if err := json.Unmarshal(bytes, &tc); err != nil {
			panic(err)
		}

		out := maxPathDP(tc)

		fmt.Println(f, out)
	}
}

// pure depth-first-search 2^n (ไม่ผ่าน hard)
func maxPath(input [][]int) int {
	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		if row == len(input) {
			return 0
		}

		if col == len(input[row]) {
			return 0
		}

		var (
			left  = dfs(row+1, col)
			right = dfs(row+1, col+1)
		)

		return input[row][col] + max(left, right)
	}

	return dfs(0, 0)
}

const (
	unseen = -1
)

// depth-first-search + dynamic programming
func maxPathDP(input [][]int) int {
	n := len(input)
	dp := make([][]int, n)

	for i := range n {
		m := len(input[i])
		dp[i] = make([]int, m)

		for j := range m {
			dp[i][j] = unseen
		}
	}

	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		if row == len(input) {
			return 0
		}

		if col == len(input[row]) {
			return 0
		}

		if dp[row][col] != unseen {
			return dp[row][col]
		}

		var (
			left  = dfs(row+1, col)
			right = dfs(row+1, col+1)
		)

		dp[row][col] = input[row][col] + max(left, right)

		return dp[row][col]
	}

	return dfs(0, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
