package main

import "fmt"

func main(){
	str := "Hello world"	

	for _, v := range str{
		fmt.Printf(" type : %T, value : %d, text : %c\n", v, v,v)
	}
}