package main

import "fmt"

const (
	MasterRoom uint8 = 1 << iota
	LivingRoom
	BathRoom
	SmallRoom
)

func SetLight(rooms, room uint8) uint8 {
	return rooms | room
}

func ResetLight(rooms, room uint8) uint8 {
	return rooms &^ room
}

func IsLightOn(rooms, room uint8) bool {
	return rooms & room != 0
}

func TurnLights(rooms uint8) {
	if IsLightOn(rooms, MasterRoom) {
		fmt.Println("turn on MasterRoom light")
	}
	if IsLightOn(rooms, LivingRoom) {
		fmt.Println("turn on LivingRoom light")
	}
	if IsLightOn(rooms, BathRoom) {
		fmt.Println("turn on BathRoom light")
	}
	if IsLightOn(rooms, SmallRoom) {
		fmt.Println("turn on SmallRoom light")
	}
}

func main() {
	var rooms uint8 = 0
	rooms = SetLight(rooms, MasterRoom)
	rooms = SetLight(rooms, LivingRoom)
	rooms = SetLight(rooms, BathRoom)
	rooms = SetLight(rooms, SmallRoom)

	rooms = ResetLight(rooms, SmallRoom)

	TurnLights(rooms)
}