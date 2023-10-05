package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	var i int = 10
	i /= 10

	var xx int = 10
	var y float64 = 30.2
	var z float64 = float64(xx) + y
	var d int = xx + int(y)

	fmt.Println(z)
	fmt.Println(d)

	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))
}
