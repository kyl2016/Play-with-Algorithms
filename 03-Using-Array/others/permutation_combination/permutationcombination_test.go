package permutationcombination

// https://leetcode.cn/problems/VvJkup/

// LCR 083. 全排列
// 中等
// 68
// 相关企业
// 给定一个不含重复数字的整数数组 nums ，返回其 所有可能的全排列 。可以 按任意顺序 返回答案。

// 示例 1：

// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 示例 2：

// 输入：nums = [0,1]
// 输出：[[0,1],[1,0]]
// 示例 3：

// 输入：nums = [1]
// 输出：[[1]]

import (
	"fmt"
	"testing"
)

func TestPermutationCombination(t *testing.T) {
	fmt.Println(PermutationCombination([]int{1, 2}))
	res := PermutationCombination([]int{1, 2, 3, 4, 5})
	fmt.Println(len(res))
}

func PermutationCombination(in []int) (out [][]int) {
	if len(in) == 1 {
		return [][]int{in}
	}

	for _, it := range in {
		for _, v := range PermutationCombination(except(in, it)) {
			vv := []int{it}
			vv = append(vv, v...)
			out = append(out, vv)
		}
	}
	return
}

func except(in []int, except int) []int {
	out := make([]int, 0)
	for _, v := range in {
		if v == except {
			continue
		}
		out = append(out, v)
	}
	return out
}
