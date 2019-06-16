package study

import (
	"fmt"
)

func DeclareArray() {
	var arr [10]int

	fmt.Printf("%v", arr)
	fmt.Print("\n")
}

func InitEmptyArray() {

	// 声明数组
	var arr = []int{}

	fmt.Printf("%v", arr)
	fmt.Print("\n")
}

func InitValueArray() {
	var arr = []int{1, 2, 3}

	fmt.Printf("%v", arr)
	fmt.Print("\n")
}

func InitLenValueArray() {
	var arr = [10]int{1, 2}
	fmt.Printf("%v", arr)
	fmt.Print("\n")
}
