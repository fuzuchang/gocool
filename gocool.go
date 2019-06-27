package main

import (
	"fmt"
	"gocool/config"
	studyGrammar "gocool/study/grammar"
	studyStructures "gocool/study/structures"
)

func main() {
	fmt.Println("today is good")
	config.GetConfig()

	// 定义数组
	var nums = []int{}

	for k, v := range nums {
		fmt.Println(k, v)
	}

	studyGrammar.DeclareArray()
	studyGrammar.InitEmptyArray()
	studyGrammar.InitValueArray()
	studyGrammar.InitLenValueArray()

	studyStructures.TestFunc()

}
