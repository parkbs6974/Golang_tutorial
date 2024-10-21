package main

import "fmt"

func main(){
	str := "Hello world"	
	arr := []rune(str)

	for i :=0; i < len(arr); i++ {
		fmt.Printf(" type : %t, value : %d, text : %c\n", arr[i],arr[i] ,arr[i])
	}
}