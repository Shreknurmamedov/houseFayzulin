package houseThings

type Thing struct {
	PeopleThing string
	Name        string
	Count       int
}

func AddThing() []Thing {
	return []Thing{
		{PeopleThing: "Ivan", Name: "Car", Count: 1},
		{PeopleThing: "Vera", Name: "Flat", Count: 3},
		{PeopleThing: "Alex", Name: "Shirt", Count: 11},
		{PeopleThing: "Alex", Name: "Toy", Count: 21},
		{PeopleThing: "Baby", Name: "Ball", Count: 2},
	}
}
