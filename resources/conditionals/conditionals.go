package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	fmt.Println(
		powR(3, 2, 10),
		powR(3, 3, 20),
	)
	fmt.Println(switchStmt())
	fmt.Println(switchWithNoCondition())
}

// Go's if statements need not be surrounded by parentheses ( )
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// Variables declared by the statement are only in scope until the end of the if
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// Variables declared inside an if short statement are also available inside any of the else blocks.
func powR(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func switchStmt() string {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		return "OS X."
	case "linux":
		return "Linux."
	default:
		// freebsd, openbsd,
		// plan9, windows...
		return os
	}
}

// switch with no condition
func switchWithNoCondition() string {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		return "Good morning!"
	case t.Hour() < 17:
		return "Good afternoon."
	default:
		return "Good evening."
	}
}
