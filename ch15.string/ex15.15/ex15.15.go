package main

import "fmt"

func main(){
	var str string = "hello Wolrd"
	var slice []byte = []byte(str)

	slice[2] = 'a'
	fmt.Println(str)
	fmt.Printf("%s\n", slice)

}