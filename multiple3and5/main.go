package main

import (
	"fmt"
	"strings"
)

// Create solution function that accept 1 parameter that be will multiple number 3 and 5 while each result of that multiplication is still lower than parameter inputed

// solution (10) => 23 = 3 + 5 + 6 + 9
// solution (20) => 78 = 3 + 5 + 6 + 9 + 10 + 12 + 15 + 18

func main() {
	number := 25

	fmt.Println(solution(number))
}

func solution(number int) string {
	var a int
	var arr []int
	for i := 1; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			a += i
			arr = append(arr, i)
		}
	}

	ets := strings.Trim(strings.Replace(fmt.Sprint(arr), " ", " + ", -1), "[]")

	res := fmt.Sprintf("%d = %s", a, ets)

	return res
}
