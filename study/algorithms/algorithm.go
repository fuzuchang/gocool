package study

import "fmt"

/*
 冒泡排序
*/
func BubbleSort(arr [5]int) {
	var i, j int
	fmt.Println("排序前：")
	for i = 0; i < 5; i++ {
		fmt.Println(arr[i])
	}

	for i = 0; i < 5; i++ {
		for j = i + 1; j < 5; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	fmt.Println("排序后：")

	for i = 0; i < 5; i++ {
		fmt.Println(arr[i])
	}

}

func InsertSort(arr [5]int) {

	var i int

	for i = 0; i < len(arr); i++ {
		fmt.Printf("%d：%d \n", i, arr[i])
	}

	fmt.Println("排序前：")

	for j := 1; j < len(arr); j++ {
		i := 0
		t := arr[j]
		for i = j - 1; i >= 0 && t < arr[i]; i-- {
			arr[i+1] = arr[i]
		}
		arr[i+1] = t
	}

	fmt.Println("排序后：")
	for i = 0; i < len(arr); i++ {
		fmt.Printf("%d：%d \n", i, arr[i])
	}

}
