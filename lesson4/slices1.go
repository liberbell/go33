package main

import (
	"fmt"
	"unsafe"
)

type Slice struct {
	ArrayPointer unsafe.Pointer
	Length       int
	Capacity     int
}

func main() {
	var stockPricesArray [3]float64 = [3]float64{89.1, 132.4, 243.5}
	var stockPricesSlice []float64 = []float64{89.1, 132.4, 243.5}

	fmt.Println(stockPricesArray)
	fmt.Println(stockPricesSlice)

	users := []string{"Name_1", "Name_2", "Name_3"}
	users = append(users, "Name_4")

	fmt.Println("Appended User Slicer: ", users)

	stockPriceSlices := make([]float64, 3, 5)
	fmt.Println("Slice with make - Length", len(stockPriceSlices))
	fmt.Println("Slice with make - Capacity", cap(stockPriceSlices))

	usersInternal := []string{"Name_1", "Name_2"}
	sliceInternals := (*Slice)(unsafe.Pointer(&usersInternal))
	fmt.Println("Slice Internals:", sliceInternals)
}
