package main

import (
	"fmt"
	"sync"
)

func main1() {
	count := 0
	var mutex sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()
			count++
			mutex.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(count)
	ch := 'a'
	fmt.Printf("%c\n", ch)
}
