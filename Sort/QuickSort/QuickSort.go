package quicksort

// QuickSort
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	innerQuickSort(arr, 0, len(arr)-1)
	return arr
}

func innerQuickSort(arr []int, l, r int) {
	if l >= r {
		return
	}

	pivot := arr[r]          // 中心点，枢轴
	i := l                   // 记录比 pivot 大的最左边的元素序号
	for j := l; j < r; j++ { // j 遍历从 l 到 r-1
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[r] = arr[r], arr[i] // 交换序号为 i 和 r 的两个元素

	innerQuickSort(arr, l, i-1)
	innerQuickSort(arr, i+1, r)
}
