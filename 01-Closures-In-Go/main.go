// Closure functions are very important for Concurrency in Go
package main

import "fmt"

func main() {
	sum := func(a, b int) int {
		// Closure function
		// This function will be visible only inside main() function
		return a + b
	}

	fmt.Printf("%d + %d = %d", 2, 2, sum(2, 2))
}
