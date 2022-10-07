package learn_sync

import (
	"fmt"
	"sync"
)

func showMsg(i int) {
	defer wp.Done()
	fmt.Printf("i:%v\n", i)
}

var wp sync.WaitGroup

func testWaitGroup() {
	for i := 0; i < 10; i++ {
		wp.Add(1)
		go showMsg(i)
	}
	wp.Wait()
	fmt.Println("end")
}
