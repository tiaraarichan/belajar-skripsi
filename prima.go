package main

import "fmt"

func main() {
	for i := 100; i < 1000; i++{
		j := 0;
	for k := 1; k <= i; k++{
			if i % k == 0{
			j++;
			}
		}
		if j==2 {
			fmt.Println(i)
		}
	}

}
