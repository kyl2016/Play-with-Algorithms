package BubbleSort

func BubbleSort2(a []int){
	for	i:=0;i<len(a)-1;i++{  		// 循环的次数
		flag := false
		for j:=0;j<len(a)-1-i;j++{  // 每次从0开始，到len(a)-1-j
			if a[j] > a[j+1]{
				flag = true
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
			}
		}
		if !flag{
			break
		}
	}
}

// 优化：如果一次循环，没有发生元素位置交换，则停止循环。