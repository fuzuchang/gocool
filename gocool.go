package main

import (
	"fmt"
	"gocool/config"
	study "gocool/study/grammar"
)

func main() {
	fmt.Println("today is good")
	config.GetConfig()

	// 定义数组
	var nums = []int{}

	for k, v := range nums {
		fmt.Println(k, v)
	}

	study.DeclareArray()
	study.InitEmptyArray()
	study.InitValueArray()
	study.InitLenValueArray()

}
