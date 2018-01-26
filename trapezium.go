package main
import "fmt"


func main(){
	var (
	a, b, n, c, f  float64
	)
	fmt.Println("masukkan nilai a :  ")
	fmt.Scanf("%f", &a)
	fmt.Println("masukkan nilai b: ")
	fmt.Scanf("%f", &b)
	fmt.Println("masukkaan nilai n")
	fmt.Scanf("%f", &n)

	h := (b-a)/n
fmt.Printf("nilai h : %f", h)

  for x := 0.; x < n+1; x++{
fx := x*h

	fmt.Printf("integrasi ke %f hasilnya %f \n", x, fx)
	if x != a && x != n {
		c = 2*fx
	} else {
	c = fx}

	fmt.Printf("iterasi ke %f hasilnya %f \n", x, c)
	f = f + c
	}

	d := f*h/2
	fmt.Printf("hasil perhitungan f adalah : %f \n", f)
	fmt.Printf("hasil perhitungan adalah : %f \n", d)
}
