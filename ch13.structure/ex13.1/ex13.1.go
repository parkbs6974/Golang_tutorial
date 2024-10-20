package main

import "fmt"

type House struct{
	Address string
	Size int
	Price float64
	Category string	
}

func main(){
	var house House
	house.Address = "seoul guri"
	house.Size = 28
	house.Price = 10
	house.Category = "apartment"

	// fmt.Println(house)
	fmt.Printf("%v\n", house)
}