package houseRooms

type Room struct {
	Name         string
	LivingPeople bool
	Count        int
}

func AddRoom() []Room {
	return []Room{
		{Name: "Kitchen", LivingPeople: true, Count: 1},
		{Name: "Bedroom", LivingPeople: false, Count: 1},
		{Name: "Bathroom", LivingPeople: false, Count: 3},
		{Name: "Loft", LivingPeople: true, Count: 2},
		{Name: "hallway", LivingPeople: true, Count: 2},
	}
}
