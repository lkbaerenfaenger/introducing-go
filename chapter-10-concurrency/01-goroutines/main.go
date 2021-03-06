package main

import "fmt"

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

// This program consists of two goroutines.
// The first goroutine is implicit and is the `main` function itself.
// The second goroutine is created when we call go `f(0)`.
func main() {
	go f(0)
	var input string
	fmt.Scanln(&input)
}
