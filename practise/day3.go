package practise

import (
	"fmt"
	"time"
)

func show() {
	for i := 0; i < 5; i++ {
		fmt.Printf("i: %d\n", i)
		time.Sleep(time.Microsecond * 300)
	}

}

func init() {
	fmt.Printf("test go \n")
	go show()
	go show()
	time.Sleep(time.Microsecond * 1000)
	gor := make(chan string, 5) // 有缓冲
	gor <- "test"
	data := <-gor
	fmt.Printf("data: %v\n", data)
}
