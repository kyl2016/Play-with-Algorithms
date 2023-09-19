package main

// 将字符串 "1234567890123" 以Z字形排列成给定的行数

import "fmt"

func main() {
	fmt.Println(resort("1234567890123", 50))
}

func resort(src string, lineCount int) string {
	if lineCount <= 0 || len(src) == 0 {
		return ""
	}
	if lineCount == 1 {
		return src
	}
	fmt.Println("src:", src, "lineCount:", lineCount)

	arr := make([][]string, lineCount)
	for i := 0; i < lineCount; i++ {
		arr[i] = make([]string, len(src))
	}
	isUpToDown := true
	row, column := 0, 0
	charIndex := 0

	first := true

	upToDown := func() {
		if first {
			for i := 0; i < lineCount; i++ {
				arr[row+i][column] = string(src[charIndex])
				if charIndex >= len(src)-1 {
					return
				}
				charIndex++
			}
			first = false
		} else {
			for i := 1; i < lineCount; i++ {
				arr[row+i][column] = string(src[charIndex])
				if charIndex >= len(src)-1 {
					return
				}
				charIndex++
			}
		}
	}

	downToUp := func() {
		for row := lineCount - 2; row >= 0; row-- {
			column++
			arr[row][column] = string(src[charIndex])
			if charIndex >= len(src)-1 {
				return
			}
			charIndex++
		}
	}

	for {
		if isUpToDown {
			upToDown()
			isUpToDown = false
		} else {
			downToUp()
			isUpToDown = true
		}

		if charIndex >= len(src)-1 {
			break
		}
	}

	result := ""
	for _, row := range arr {
		for _, col := range row {
			if col != "" {
				result += col
				fmt.Print(col)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
	return result
}
