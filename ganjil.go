package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {
		if i % 2 != 1 {
			continue
		}

		if i > 19 {
			break
		}

	fmt.Println("Ganjil", i)
	}
}
