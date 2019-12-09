package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex

func main() {
	var a [10]int
	fmt.Println(a) //{0,0,0,0,..}
	//lock.Lock()
	for i := 0; i < 10; i++ {
		go func(ii int) { //race condition!
			for {
				a[ii]++
				//runtime.Gosched()
			}
		}(i)
	}
	//lock.Unlock()
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
