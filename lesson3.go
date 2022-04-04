package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	//Simple For
	fmt.Println(CycleFor())
	//Short For
	fmt.Println(CycleForShortLikeWhile())
	//For is GO's "while"
	fmt.Println(CycleWhileInGoFor())
	//Forever
	//Forever()
	//Test if
	fmt.Println(IfTest(9), IfTest(-4))
	//Short IF
	fmt.Println(IfShortStatement(3, 2, 10))
	fmt.Println(IfShortStatement(3, 3, 20))
	//If and ELSE
	fmt.Println(
		IfAndElse(3, 2, 10),
		IfAndElse(3, 3, 20),
	)
	//LOOPS AND FUNCTIONS
	fmt.Println("result = ", Sqrt(121))
	//SWITCH
	fmt.Println(SwitchTest())
	//SWITCH with no condition
	fmt.Println(SwithcWithNoCondition())
	//Defer
	deferHelloWorld()
	//Stacking defer
	deferStacking()
}

func deferStacking() {
	fmt.Println("counting")
	for i:=0;i<10;i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func deferHelloWorld() {
	defer fmt.Println("go's world")

	fmt.Print("Hello ")
}

func SwithcWithNoCondition() string {
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

func SwitchTest() string {
	fmt.Print("Go runs on ")
	result := "Go runs on "
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
		result += "OS X."
	case "linux":
		fmt.Println("Linux.")
		result += "Linux."
	default:
		fmt.Printf("%s.\n", os)
		result += os
	}
	return result
}

func Sqrt(x float64) float64 {
	var z = float64(1)
	for i := 1; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z

}

func IfAndElse(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func IfShortStatement(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func CycleFor() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

func CycleForShortLikeWhile() int {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	return sum
}

func CycleWhileInGoFor() int {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	return sum
}

func Forever() {
	for {
	}
}

func IfTest(x float64) string {
	if x < 0 {
		return IfTest(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
