package main

import (
	"fmt"
	"os"
	"strconv"
)

func UNUSED(x ...interface{}) {}

func calc(operator int, values []int) int {
	sum := 0
	switch operator {
	case 1:
		for i := 0; i < len(values); i++ {
			sum += int(values[i])
			if i == len(values)-1 {
				fmt.Print(values[i], " = ")
			} else {
				fmt.Print(values[i], " + ")
			}
		}
		return sum

	case 2:
		sum += values[0] * 2
		for i := 0; i < len(values); i++ {
			sum -= int(values[i])
			if i == len(values)-1 {
				fmt.Print(values[i], " = ")
			} else {
				fmt.Print(values[i], " - ")
			}
		}
		return sum

	case 3:
		sum = 1
		for i := 0; i < len(values); i++ {
			sum *= int(values[i])
			if i == len(values)-1 {
				fmt.Print(values[i], " = ")
			} else {
				fmt.Print(values[i], " * ")
			}
		}
		return sum
	default:
		return 0
	}
}

func main() {

	var vals = make([]int, len(os.Args)-2)

	for i := 2; i < len(os.Args); i++ {
		num, err := strconv.ParseInt(os.Args[i], 10, 0)
		vals[i-2] = int(num)
		UNUSED(err)
	}

	if os.Args[1] == "add" {
		fmt.Println(calc(1, vals))

	} else if os.Args[1] == "sub" {
		fmt.Println(calc(2, vals))

	} else if os.Args[1] == "mult" {
		fmt.Println(calc(3, vals))

	} else {
		fmt.Println("invalid input")
	}

}
