package main

import (
	"fmt"
	"unsafe"
)

type User struct{
	Age  int
	Score float64
}


func main(){
	var a int =40
	user := User{23, 44.5}
	fmt.Println(unsafe.Sizeof(user), unsafe.Sizeof(a))
}