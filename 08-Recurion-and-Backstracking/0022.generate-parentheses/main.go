package main

import "fmt"

func main() {
	r := generateParenthesis(3)
	fmt.Println(r)
}

var result []string
func generateParenthesis(n int) []string {
	result = nil
	generate2(0, 0, n, "")
	//generate(0, 2*n, "")
	return result
}

func generate(level, max int, s string) {
	// terminator
	if level >= max {
		// filter the invalid s

		fmt.Println(s)
		return
	}
	// process current logic:
	// drill down
	generate(level+1, max, s+"(")
	generate(level+1, max, s+")")
	// reverse states
}

// left 随时可以加，只要别超标
// right 必须出现在 left 之后，且左个数>右个数
// 做了剪枝
func generate2(left, right, n int, s string) {
	if left == n && right == n {
		//fmt.Println(s)
		result = append(result, s)
		return
	}
	if left < n {
		generate2(left+1, right, n, s+"(")
	}
	if left > right {
		generate2(left, right+1, n, s+")")
	}
}