package main

import "fmt"

func main(){
	var a []int = []int{1,2,-2,4,6,-7,1,-9,3,-1,8,0}
	fmt.Println(a)

	BubbleSort(a)
	fmt.Println(a)
}

func BubbleSort(a []int){
	for	i:=0;i<len(a)-1;i++{  		// 循环的次数
		for j:=0;j<len(a)-1-i;j++{  // 每次从0开始，到len(a)-1-j
			if a[j] > a[j+1]{
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
			}
		}
	}
}

// 优化：如果一次循环，没有发生元素位置交换，则停止循环。