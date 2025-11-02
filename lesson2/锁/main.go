package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) jia() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
func main() {
	counter := Counter{}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(10)
		go func() {
			for j := 0; j < 10; j++ {
				counter.jia()
				wg.Done()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.Value())
}
