package main

import (
	"fmt"
	"strings"
)

// Create a function that receive input and will give output:
// Reverse each word of these:

// "I am A Great human"
// into
// "I ma A Taerg namuh"

// note: observe the Capital word behavior. Capital only when in the first letter

func main() {
	str := "I am A Great human"
	fmt.Println(reverseWords(str))
}

func reverseWords(str string) string {
	reverse := func(str string) string {
		var res string
		for _, v := range str {
			res = string(v) + res
		}
		return res
	}

	arrStr := strings.Split(str, " ")
	for idx, val := range arrStr {
		strReverse := reverse(val)
		if val == strings.Title(val) {
			arrStr[idx] = strings.Title(strings.ToLower(strReverse))
			continue
		}
		arrStr[idx] = strReverse
	}
	return strings.Join(arrStr, " ")
}
