package main

import "fmt"

func main(){
	var a []int = []int{1,2,-2,4,6,-7,1,-9,3,-1,8,0}
	fmt.Println(a)

	InsertionSort(a)
	fmt.Println(a)
}

func InsertionSort(a []int) {
	for i:=1;i<len(a);i++{
		for j:=i;j>0;j--{
			if a[j]<a[j-1]{
				a[j],a[j-1] = a[j-1],a[j]
			}
		}
	}
}