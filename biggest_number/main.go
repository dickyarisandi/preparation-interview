package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Create a function that takes one positive three digit integer and rearranges its digits to get the maximum possible number. Assume that the argument is an integer.

// Returm null if the argument is invalid.

// maxRedigit(123) --> 321
// maxRedigit(231) --> 321
// maxRedigit(321) --> 321

// maxRedigit(-1) --> null
// maxRedigit(0) --> null
// maxRedigit(99) --> null
// maxRedigit(1000) --> null

func main() {
	res := maxRedigit(1234)
	fmt.Println(*res)
}

func maxRedigit(data int) *int {
	if data <= 0 {
		return nil
	}

	strData := strconv.Itoa(data)
	arrData := strings.Split(strData, "")

	a := 0
	sort.Slice(arrData, func(i, j int) bool {
		if arrData[i] > arrData[j] {
			a++
		}
		return arrData[i] > arrData[j]
	})

	if a < 1 {
		return nil
	}

	dataStr := strings.Join(arrData, "")
	res, _ := strconv.Atoi(dataStr)

	return &res
}
