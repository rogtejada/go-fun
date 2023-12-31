//simple echo command with two approaches

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[:], " "))

	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}
