package main

import (
	"actors/actor"
	"fmt"
	"math"
	"time"
)

func main() {
	const (
		EventAdd = iota
		EventMultiply
		EventSquare
	)

	// Create an actor
	a := actor.New()

	// Register event handlers
	a.Register(EventAdd, func(args ...actor.Arg) {
		i1 := args[0].(*int)
		i2 := args[1].(int)

		*i1 += i2
	})

	a.Register(EventMultiply, func(args ...actor.Arg) {
		i1 := args[0].(*int)
		i2 := args[1].(int)

		*i1 *= i2
	})

	a.Register(EventSquare, func(args ...actor.Arg) {
		i1 := args[0].(*int)

		*i1 = int(math.Pow(float64(*i1), 2))
	})

	x := 2

	a.Cast(EventAdd, &x, 1)
	a.Cast(EventMultiply, &x, 10)
	a.Cast(EventSquare, &x)

	<-time.After(time.Second)
	fmt.Println("x =", x)

	// Close the actor
	a.Close()
}
