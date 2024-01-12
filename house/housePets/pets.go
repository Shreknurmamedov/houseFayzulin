package housePets

type Pet struct {
	PeoplePet string
	Name      string
}

func AddPet() []Pet {
	return []Pet{
		{PeoplePet: "Vera", Name: "Cat"},
		{PeoplePet: "Vera", Name: "Cat"},
		{PeoplePet: "Ivan", Name: "Dog"},
		{PeoplePet: "Ivan", Name: "Dog"},
		{PeoplePet: "Max", Name: "Rabbit"},
	}
}
