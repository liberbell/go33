package main

import "fmt"

type Point struct {
	X int
	Y int
}

func modify(x *int) {
	*x = 50
}

func (p *Point) changeCordinates(xNew int, yNew int) {
	p.X = xNew
	p.Y = yNew
}

func main() {

	var n = 10
	var p = &n
	fmt.Printf("Type of Pointer: %T\n", p)

	fmt.Printf("Address: %v\n", p)

	fmt.Printf("Value: %v\n", *p)

	var a int = 50
	var b *int = &a

	fmt.Println(a)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)
	fmt.Println(*b)

	var c int = 10

	fmt.Printf("Original value: %d\n", c)
	modify(&c)

	fmt.Printf("Modified value: %d\n", c)

	p := &Point{X: 5, Y: 3}
	fmt.Println(*p)

}
