package main

import (
	"fmt"
	"regexp"
	"strings"
)

func findFirstStringInBracket(str string) string {
	var element string
	if (len(str)) > 0 {
		re := regexp.MustCompile(`\((.*?)\)`)
		fmt.Printf("Pattern: %v\n", re.String()) // print pattern

		submatchall := re.FindAllString(str, -1)
		element = strings.Trim(submatchall[0], "(") //submatchall[0] is first element of array
		element = strings.Trim(element, ")")
	} else {
		return ""
	}

	return element
}

func main() {
	line_test := "(hehe) (haha)"
	result := findFirstStringInBracket(line_test)
	fmt.Println(result)

	line_test = "hehe) (haha)"
	result = findFirstStringInBracket(line_test)
	fmt.Println(result)

}