package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius int
}

func main() {
	fmt.Println("The area of the Circle is", areaOfACircle())
	fmt.Println("The circumference of the Circle is", circumferenceOfACircle())
}

func areaOfACircle() float64 {
	circle := Circle{7}

	return 2 * math.Pi * float64(circle.radius)
}

func circumferenceOfACircle() float64 {
	circle := Circle{7}

	return math.Pi * math.Pow(float64(circle.radius), 2)
}
