package main

import "fmt"

func main(){
	const pi1 float64 = 3.14159265468989228 
	var pi2 float64 = 3.14159265468989228 

	// pi1 = 3
	pi2 = 4

	fmt.Printf("radius : %f\n", pi1)
	fmt.Printf("radius : %f\n", pi2)

}