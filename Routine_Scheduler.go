package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(3)
	var processTest sync.WaitGroup
	processTest.Add(3)
	go func() {
		defer processTest.Done()
		for i := 0; i < 5; i++ {
			for j := 1; j <= 5; j++ {
				fmt.Printf(" %d", j)
				if j == 3 {
					fmt.Println()
				}
			}
		}
	}()
	go func() {
		defer processTest.Done()
		for j := 0; j < 5; j++ {
			for char := 'A'; char < 'A'+4; char++ {
				fmt.Printf("%c ", char)
				if char == 'C' {
					fmt.Println()
				}
			}
		}
	}()

	go func() {
		defer processTest.Done()
		for i := 0; i < 5; i++ {
			for j := 0; j <= 3; j++ {
				fmt.Printf(" %d", j)
				if j == 5 {
					fmt.Println()
				}
			}
		}
	}()
	processTest.Wait()
}
