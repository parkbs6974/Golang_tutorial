package main

import "fmt"

func main(){
	temp := 3

	if temp > 28 {
		fmt.Println("turn on a/c")
		} else if temp < 3{
			fmt.Println("turn on heater")
			} else if temp <= 18 {
				fmt.Println("let's go out")
	} else {
		fmt.Println("hot")
	}
}