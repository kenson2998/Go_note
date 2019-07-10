package main //main套件
import ( //import fmt,內建的標準庫  ,多的時候也可以用import()
	"fmt" //fmt是標準庫,但fmt不知道是哪裡來的
	"sync"
)

var wg sync.WaitGroup

func print(a int) {
	fmt.Println("Hello", a)
	wg.Done() //每Done一次就代表他做完

}

func main() {
	wg.Add(4) // 有幾個thread 就給他等待的整數
	go print(1)
	go print(2)
	go print(3)
	go print(4)
	wg.Wait() // 等到add裡面變成0,就會離開

}
