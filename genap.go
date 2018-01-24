package main

import "fmt"

func main(){
	fmt.Println("Menampilkan bilangan genap")
	

	for i := 1; i <= 20; i++{
		if i % 2 == 1 {
			continue
		}
		if i > 20 {
			break
		}

	fmt.Println("Genap", i)

	}
}
