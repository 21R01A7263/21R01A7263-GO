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
		for i := 0; i < 3; i++ {
			for j := 5; j <= 10; j++ {
				fmt.Printf(" %d", j)
				if j == 7 {
					fmt.Println()
				}
			}
		}
	}()

	go func() {
		defer processTest.Done()
		for j := 0; j < 10; j++ {
			for char := 'A'; char < 'A'+10; char++ {
				fmt.Printf("%c ", char)
				if char == 'F' {
					fmt.Println()
				}
			}
		}
	}()

	processTest.Wait()
}
