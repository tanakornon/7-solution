package main

import (
	"fmt"
	"strings"
)

type TestCase struct {
	Input  string
	Output string
}

var result = map[bool]string{
	true:  "PASS",
	false: "FAIL",
}

func main() {
	tcs := []TestCase{
		{
			Input:  "LLRR=",
			Output: "210122",
		},
		{
			Input:  "==RLL",
			Output: "000210",
		},
		{
			Input:  "=LLRR",
			Output: "221012",
		},
		{
			Input:  "RRL=R",
			Output: "012001",
		},
	}

	for i, tc := range tcs {
		expect := tc.Output
		actual := decodeV2(tc.Input)

		fmt.Printf("CASE #%d input: %s expect: %s actual: %s --- %s\n",
			i+1,
			tc.Input,
			expect,
			actual,
			result[actual == expect],
		)
	}

}

const (
	L byte = 'L'
	R byte = 'R'
	E byte = '='
)

// v1
func decodeV1(str string) string {
	var (
		n   = len(str)
		num = make([]int, n+1)
	)

	for i := range str {
		switch str[i] {
		case R:
			num[i+1] = num[i] + 1
		case E:
			num[i+1] = num[i]
		}

		if str[i] == L && (i+1 == n || str[i+1] != L) {
			for j := i; j >= 0; j-- {

				switch str[j] {
				case L:
					if num[j] <= num[j+1] {
						num[j] = num[j+1] + 1
					}
				case E:
					num[j] = num[j+1]
				case R:
					break
				}

			}
		}
	}

	var b strings.Builder
	for _, v := range num {
		b.WriteString(fmt.Sprint(v))
	}

	return b.String()
}

// v2 refactored
func decodeV2(str string) string {
	var (
		n   = len(str)
		num = make([]int, n+1)
	)

	for i := range n {
		switch str[i] {
		case R:
			num[i+1] = num[i] + 1
		case E:
			num[i+1] = num[i]
		}
	}

	for j := range n {
		i := n - j - 1

		switch str[i] {
		case L:
			num[i] = max(num[i], num[i+1]+1)
		case E:
			num[i] = num[i+1]
		}
	}

	var b strings.Builder
	for _, v := range num {
		b.WriteString(fmt.Sprint(v))
	}

	return b.String()
}

// --- กระดาษทด ---

//  L L R R =
// 2 1 0 1 2 2

//  L L R R
// 2 1 0 1 2 0

//  = = R L L
// 0 0 0 2 1 0

//  = L L R R
// 2 2 1 0 1 2

//  R R L = R
// 0 1 2 0 0 1

const (
	ZeroASCII = 48
)

func operation(v int, r byte) (int, int) {
	switch r {
	case L:
		return v + 1, v
	case R:
		return v, v + 1
	default:
		return v, v
	}
}
