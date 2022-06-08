package main

import (
	"fmt"
	"strings"
)

// Alternate each case of each of string given

// alternateCase("abc") => "ABC"
// alternateCase("ABC") => "abc"
// alternateCase("Hello World") => "hELLO wORLD"

func main() {
	fmt.Println(alternateCase("Hello World"))
}

func alternateCase(s string) string {
	arrS := strings.Split(s, "")

	var abj []string
	for _, v := range arrS {
		if v == strings.ToLower(v) {
			abj = append(abj, strings.ToUpper(v))
			continue
		}

		abj = append(abj, strings.ToLower(v))
	}

	result := strings.Join(abj, "")

	return result
}
