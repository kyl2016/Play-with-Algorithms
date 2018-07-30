package main

import "fmt"

func main(){
	var a []int = []int{1,2,-2,4,6,-7,1,-9,3,-1,8,0}
	fmt.Println(a)

	QuickSort(a,0,len(a)-1)
	fmt.Println(a)
}

func QuickSort(a []int, l int, r int){
	if l >= r {
		return
	}

	var k int = l+1
	for i:=l+1;i<=r;i++{
		if a[i] <= a[l]{
			
			a[i],a[k] = a[k],a[i]
			k++
		}
	}

	a[l],a[k-1] = a[k-1],a[l]

	fmt.Println(a)

	QuickSort(a, l, k-2)
	QuickSort(a, k, r)
}