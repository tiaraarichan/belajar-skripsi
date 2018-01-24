package main

import "fmt"


func main() {
	fmt.Println("Menampilkan bilangan")

	for i := 0 ; i<5; i++{
		j := i*i
		fmt.Println((i), "" ,"dikali", i, "", "adalah", j)
}
}
