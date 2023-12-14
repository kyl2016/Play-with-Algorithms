package removeelement

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	removeElement([]int{3, 2, 2, 3}, 3)
	removeElement([]int{3, 2, 2, 3}, 2)
	removeElement([]int{3, 2, 2, 3, 5}, 5)
	removeElement([]int{5, 3, 2, 2, 3}, 5)
}

func removeElement(nums []int, val int) int {
	lastEqualIndex := 0
	for i, it := range nums {
		if it != val {
			if lastEqualIndex < i {
				nums[lastEqualIndex] = it
			}
			lastEqualIndex += 1
		}
	}

	fmt.Println(nums[:lastEqualIndex])

	return lastEqualIndex
}
