package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a float64, init_v float64, init_s float64) func(t float64) float64 {
	var fn = func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + init_v*t + init_s
	}
	return fn
}

func main() {
	var a float64
	var init_v float64
	var init_s float64
	var time float64

	fmt.Print("Enter acceleration value (a): ")
	fmt.Scanln(&a)
	fmt.Print("Enter initial velociry (v_0): ")
	fmt.Scanln(&init_v)
	fmt.Print("Enter initial displacement (s_0): ")
	fmt.Scanln(&init_s)
	var genFn = GenDisplaceFn(a, init_v, init_s)

	fmt.Print("Enter time (s): ")
	fmt.Scanln(&time)
	fmt.Printf("Displacement value after %fs: %f", time, genFn(time))
}
