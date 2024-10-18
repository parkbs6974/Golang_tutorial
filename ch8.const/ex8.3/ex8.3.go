package main

import "fmt"

const Pig int = 0
const Cow int = 1
const Chicken int =2

func PrintAnimal(animal int){
	if animal == Pig{
		fmt.Println("pigggg")
		} else if animal == Cow{
		fmt.Println("cowwww")
		} else if animal == Chicken{
		fmt.Println("chickennnnn")
	} else {
		fmt.Println("....")
	}
}

func main(){
	PrintAnimal(Pig)
	PrintAnimal(Cow)
	PrintAnimal(Chicken)
}