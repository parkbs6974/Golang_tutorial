package main

import "fmt"

func main(){
	var t [5]float64 = [5]float64{24.0, 25.9, 39.2, 39.1, 39.5}

	for i, v := range t{
		fmt.Println(i, v)
	}
}