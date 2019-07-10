package main

import (
	"fmt"
)

func main() {

	var nums [5]int
	nums = [5]int{3, 2, 7, 11, 3}
	//target := 6

	for index,_ := range nums {
		for index2 := index+1  ; index2 < len(nums) ; index2++ {
			fmt.Println(index,index2)
		}
	}
}
