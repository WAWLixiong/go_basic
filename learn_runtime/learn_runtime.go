package learn_runtime

import (
	"fmt"
	"runtime"
	"time"
)

func show() {
	for i := 0; i < 5; i++ {
		fmt.Printf("sub %d\n", i)
		if i > 5 {
			runtime.Goexit() // 只退出当前所在的goroutine
		}
	}
}

func testRuntimeGoSched() {
	go show()

	time.Sleep(time.Second * 2)

	for i := 0; i < 2; i++ {
		// runtime.Gosched() // 让出控制权让其他 goroutine先运行
		fmt.Printf("main %d\n", i)

	}
}
