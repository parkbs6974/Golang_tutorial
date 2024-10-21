package main

import (
	"fmt"
	"strings"
)

func ToUpper(str string) string {
var builder strings.Builder

	for _, c := range str {
		if c > 'a' && c <= 'z' {
			builder.WriteRune('A' + (c - 'a'))
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}

func main(){
	var str string = "hello Wolrd"
	var str1 string = "test Using"
	fmt.Println(ToUpper(str))
	fmt.Println(strings.ToUpper(str1))


}