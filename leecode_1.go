package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		for ii, vv := range nums {

			switch v+vv == target && i != ii {
			case true:
				return []int{i, ii}
			}

		}
	}
	return []int{}
}

func main() {
	a := []int{2, 7, 11, 15}
	t := 9
	fmt.Println(twoSum(a, t))
}

/// go routine 寫法
//package main
//import (
//	"fmt"
//	"sync"
//)
//
//var wg sync.WaitGroup
//var na []int
//
//func second(a []int, b []int, target int) []int {
//	for i, v := range b {
//		if a[1]+v == target {
//			na = []int{a[0], i + 1}
//		}
//	}
//	wg.Done()
//	return []int{}
//}
//func twoSum(nums []int, target int) []int {
//	wg.Add(len(nums))
//	for i, v := range nums {
//		go second([]int{i, v}, nums[i+1:], target)
//	}
//	wg.Wait()
//	if len(na) >1{return na}
//	return []int{}
//}
//func main() {
//	a := []int{2, 7, 11, 15}
//	t := 9
//
//	twoSum(a, t)
//
//	fmt.Println(na)
//}
