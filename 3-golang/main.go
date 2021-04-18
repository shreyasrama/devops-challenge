package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(a int) {
			fmt.Println(a)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
