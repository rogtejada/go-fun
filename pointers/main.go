package main

import "fmt"

//Car struct
type Car struct {
	Model string
	Price float32
	Year  int
}

func main() {
	a := Car{
		Model: "corsa",
		Price: 100.32,
		Year:  1993,
	}

	//print corsa
	fmt.Println(a)

	//call function that mutates car model
	b := changeCarModel(a)

	//a object remains the same
	fmt.Println(a)
	//only returned has new model
	fmt.Println(b)

	//call function passing pointer to a object
	c := changeCarModelPointer(&a)

	//a object gets changed
	fmt.Println(a)

	//c is a pointer to A
	fmt.Println(c)

	a.Model = "fusion"

	//both are changed
	fmt.Println(a)
	fmt.Println(c)
}

func changeCarModel(c Car) Car {
	c.Model = "fusca"

	return c
}

func changeCarModelPointer(c *Car) *Car {
	c.Model = "fusca"

	return c
}
