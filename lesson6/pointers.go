package main

import "fmt"

type Point struct {
	X int
	Y int
}

func modify(x *int) {
	*x = 50
}

func (p *Point) ChangeCordinates(xNew int, yNew int) {
	p.X = xNew
	p.Y = yNew
}

func modifySlice(slicer *[]int) {
	*slicer = append(*slicer, 100)
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

	q := &Point{X: 5, Y: 3}
	fmt.Println(*q)

	q.ChangeCordinates(1, 2)
	fmt.Println(*q)

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Original Slice: ", numbers)
	modifySlice(&numbers)
}
