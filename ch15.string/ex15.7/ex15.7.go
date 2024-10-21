package main

import "fmt"

func main(){
	str := "hello world"

	for i := 0; i < len(str); i++ {
		fmt.Printf("type : %t value : %d letter : %c\n", str[i],str[i],str[i])
	}
}