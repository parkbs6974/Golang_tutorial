package main

import "fmt"

func HasRichFriend()bool{
	return true
}

func GetFriendsCount()int{
	return 3
}

func main(){
	price := 35000

	if price >50000{
		if HasRichFriend(){
			fmt.Println("need to tie my shoose")
		} else{
			fmt.Println("need to split bill")
		}
	} else if price >= 30000 {
		if GetFriendsCount() > 3 {
			fmt.Println("shoooot my shoooose")
		} else {
			fmt.Println("let's split bill, not many poeple")
		}
		} else {
		fmt.Println("it's on me")
	}
}