package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	stdin := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("insert number")
		var number int 
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println("please insert number")

			stdin.ReadString('\n')
			continue
		}

		fmt.Printf("the number of what you just put %d.\n", number)
		if number % 2 == 0 {
			break
		}
	}
	fmt.Println("the loop is closed")
}