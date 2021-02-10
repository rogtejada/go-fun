package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func main() {
	word := flag.String("word", "foo", "a string")
	num := flag.Int("int", 32, "a number")
	bol := flag.Bool("bol", false, "a boolean")

	//run a command in bash
	cmd := exec.Command("bash", "-c", "echo output from echo")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Print(string(stdout))


	//execute command flag parsing
	flag.Parse()

	//need to dereference the pointers to get the actual option values
	fmt.Println("word:", *word)
	fmt.Println("num:", *num)
	fmt.Println("bol:", *bol)
	fmt.Println(flag.Args())
}