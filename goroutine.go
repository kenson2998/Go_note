package main //main套件
import ( //import fmt,內建的標準庫  ,多的時候也可以用import()
	"fmt" //fmt是標準庫,但fmt不知道是哪裡來的
	"sync"
)

var wg sync.WaitGroup
var na []int

func second(a []int, b []int, target int) []int {
	//fmt.Println(a, b)
	for i, v := range b {
		if a[1]+v == target {
			na = []int{a[0], i + 1}
		}
	}
	wg.Done()
	return []int{}
}
func twoSum(nums []int, target int) []int {
	wg.Add(len(nums)) // 有幾個thread 就給他等待的整數
	for i, v := range nums {
		go second([]int{i, v}, nums[i+1:], target)
	}
	wg.Wait()
	if len(na) >1{return na}
	return []int{}
}
func main() {
	a := []int{2, 7, 11, 15}
	t := 9

	twoSum(a, t)

	// 等到add裡面變成0,就會離開
	fmt.Println(na)
}
