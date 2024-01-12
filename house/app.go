package main

import (
	"houseFayzulin/house/house"
)

func main() {
	newHouse := house.CreateHouse()
	house.ShowHouse(newHouse)
}
