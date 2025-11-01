package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func download(filename string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Printf("%s 下载完成", filename)
}

func main() {
	filename := []string{"file1.zip", "file2.pdf", "file3.mp4"}

	results := make(chan string, 3)
	var wg sync.WaitGroup
	for _, na := range filename {
		wg.Add(1)
		go download(na, &wg, results)
	}
	wg.Wait()
	fmt.Println("所有文件下载完成!")
}
