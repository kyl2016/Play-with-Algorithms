package main

import "fmt"

func main(){
	var a []int = []int{1,2,-2,4,6,-7,1,-9,3,-1,8,0}
	fmt.Println(a)

 	r := MergeSort(a, 0, len(a)-1)
	fmt.Println(r)
}

func MergeSort(a []int, l int, r int) []int{
	if l>=r {
		return a[l:l+1]
	}

	mid := l + (r-l)/2;

	left := MergeSort(a, l, mid)
	right := MergeSort(a, mid+1, r)

	return Combine(left, right)
}

func Combine(left []int, right []int) []int{
	var result []int = make([]int, len(left)+len(right))

	i := 0
	l,r := 0,0
	for l < len(left) && r < len(right){
		if left[l] < right[r] {
			result[i] = left[l]
			l++
		}else {
			result[i] = right[r]
			r++
		}

		i++
	}

	for l < len(left) {
		result[i] = left[l]
		l++
		i++
	}

	for r < len(right){
		result[i] = right[r]
		r++
		i++
	}

	return result
}