package house

import (
	"fmt"
	"houseFayzulin/house/houseFamily"
	"houseFayzulin/house/houseFurniture"
	"houseFayzulin/house/housePets"
	"houseFayzulin/house/houseRooms"
	"houseFayzulin/house/houseThings"
)

type House struct {
	Length       float64
	Height       float64
	NewFamily    []houseFamily.Family
	NewRooms     []houseRooms.Room
	NewFurniture []houseFurniture.Furniture
	NewThings    []houseThings.Thing
	NewPets      []housePets.Pet
}

func CreateHouse() House {
	return House{
		Length:       4.567,
		Height:       8.999,
		NewFamily:    houseFamily.AddMember(),
		NewRooms:     houseRooms.AddRoom(),
		NewFurniture: houseFurniture.AddFurniture(),
		NewThings:    houseThings.AddThing(),
		NewPets:      housePets.AddPet(),
	}
}

func ShowHouse(house House) {
	fmt.Printf("House:\n")
	fmt.Printf("Length: %0.3f\n", house.Length)
	fmt.Printf("Height: %0.3f\n", house.Height)

	showFamily(house.NewFamily)
	showRooms(house.NewRooms)
	showFurniture(house.NewFurniture)
	showThings(house.NewThings)
	showPets(house.NewPets)
}

func showFamily(houseObject []houseFamily.Family) {
	for _, object := range houseObject {
		fmt.Printf("- %s %s\n", object.Member, object.Name)
	}
}

func showRooms(houseObjects []houseRooms.Room) {
	for _, object := range houseObjects {
		fmt.Printf("- %s: living people: %t, count: %d\n", object.Name, object.LivingPeople, object.Count)
	}
}

func showFurniture(houseObjects []houseFurniture.Furniture) {
	for _, object := range houseObjects {
		fmt.Printf("- %s: %.2f m\n", object.Name, object.Size)
	}
}

func showThings(houseObjects []houseThings.Thing) {
	for _, object := range houseObjects {
		fmt.Printf("- %s: people thing: %s, count: %d\n", object.PeopleThing, object.Name, object.Count)
	}
}

func showPets(houseObjects []housePets.Pet) {
	for _, object := range houseObjects {
		fmt.Printf("- People pet: %s, pet: %s\n", object.PeoplePet, object.Name)
	}
}
