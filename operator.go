package main
import "fmt"

func main() {
var a, b, c, d int

a = 2
b = 4

c = a + b

d = 2*c

fmt.Println("a = 2")
fmt.Println("b = 4")
fmt.Println("c = a + b")
fmt.Printf("c = %d\n", c)
fmt.Println("d = 2*(a + b)")
fmt.Printf("d = %d\n", d)
}
