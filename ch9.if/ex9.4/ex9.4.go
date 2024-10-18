package main

import "fmt"

var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}


func main(){
	if false || IncreaseAndReturn() < 5 {
		fmt.Println("increase 1")
	}
	fmt.Println(cnt)
}