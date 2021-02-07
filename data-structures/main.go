package main

import "fmt"

func main() {
	mapExample()
	sliceExample()
}

func mapExample() {
	fmt.Println("Map:")
	m := make(map[string]int)

	m["a"] = 10
	m["b"] = 20

	fmt.Println(m)

	fmt.Println(m["a"])

	delete(m, "a")

	fmt.Println(m)

	x, isPresent := m["a"]
	fmt.Println("x:", x, "isPresent:", isPresent)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)
}

func sliceExample() {
	fmt.Println("Slices:")
	s := make([]string, 3)
	fmt.Println(s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)
	fmt.Println(s[2])

	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))

	//append beyond slice cap
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	//create another slice from the previous one
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	//sub array from pos 2(included) to pos 5(excluded)
	l := s[2:5]
	fmt.Println(l)

	//from start to pos 5(excluded)
	l = s[:5]
	fmt.Println(l)

	//from 2(included) to end
	l = s[2:]
	fmt.Println(l)

	t := []string{"g", "h", "i"}
	fmt.Println(t)

	//create matrix specifiyng y dimension
	matrix := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1

		//create the inner slices specifying the x dimension
		matrix[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			matrix[i][j] = i + j
		}
	}

	fmt.Println("matrix: ", matrix)
}
