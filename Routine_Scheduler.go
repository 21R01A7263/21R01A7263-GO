package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(4)
	var processTest sync.WaitGroup
	processTest.Add(2)
	go func() {
		defer processTest.Done()
		for i := 0; i < 30; i++ {
			for j := 15; j <= 25; j++ {
				fmt.Printf(" %d", j)
				if j == 18 {
					fmt.Println()
				}
			}
		}
	}()
	go func() {
		defer processTest.Done()
		for j := 0; j < 10; j++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
				if char == 'F' {
					fmt.Println()
				}
			}
		}
	}()

	processTest.Wait()
}
