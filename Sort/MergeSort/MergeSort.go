package mergeSort

func Sort(arr []int) []int {
	if len(arr) <= 1 { // 检查输入
		return arr
	}

	return mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, l, r int) []int {
	// 递归终止条件
	if l >= r {
		return arr[r : r+1]  // 注意：不是返回 arr，arr 是全部的数据，而 l 和 r 的区间只是一部分数据
	}

	mid := l + (r-l)/2 // 注意不要溢出
	left := mergeSort(arr, l, mid)
	right := mergeSort(arr, mid+1, r)
	return combine(left, right)
}

func combine(left, right []int) []int {
	var r []int

	lIndex := 0
	rIndex := 0
	for lIndex < len(left) && rIndex < len(right) { // 必须消费完其中一个数组
		if left[lIndex] < right[rIndex] {
			r = append(r, left[lIndex])
			lIndex++
		} else {
			r = append(r, right[rIndex])
			rIndex++
		}
	}

	for ; lIndex < len(left); lIndex++ { // lIndex 也没有 append，所有从 lIndex 开始比较
		r = append(r, left[lIndex])
	}

	for ; rIndex < len(right); rIndex++ {
		r = append(r, right[rIndex])
	}

	return r
}
