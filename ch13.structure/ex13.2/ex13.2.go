package main

import "fmt"

type User struct {
	Name string
	ID string
	Age int
}

type VIPUser struct {
	// UserInfo User
	User
	VIPLevel int
	Price int
}


func main(){
	user := User{"kim", "park123", 30}
	vip := VIPUser{
		User{"alice", "mango", 40},
		3,
		250,
	}

	fmt.Printf("user : %s ID: %s age : %d\n", user.Name, user.ID, user.Age)
	// fmt.Printf("VIP user %s ID : %s age : %d VIP LEVEL : %d VIP price : %d",
	// 					vip.UserInfo.Name, vip.UserInfo.ID, vip.UserInfo.Age, vip.VIPLevel, vip.Price,
	// )	
	fmt.Printf("VIP user %s ID : %s age : %d VIP LEVEL : %d VIP price : %d",
						vip.Name, vip.ID, vip.Age, vip.VIPLevel, vip.Price,
	)	 // embedded 
}