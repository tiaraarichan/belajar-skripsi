package main

import "fmt"

func main() {
 	var (
		s, hasil float64
	)

	fmt.Println("s :")
	fmt.Scanf("%f", &s)
	hasil = s * s
	fmt.Println("Luas persegi :", hasil, "m2")
}
