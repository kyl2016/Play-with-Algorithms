package main

import "fmt"

func main(){
	var a []int = []int{1,2,-2,4,6,-7,1,-9,3,-1,8,0}
	fmt.Println(a)

	BubbleSort(a)
	fmt.Println(a)
}

func BubbleSort(a []int){

	for i := 0; i<len(a)-1;i++{
		for j := i+1; j< len(a); j++{
			if a[j] < a[i]{
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
}