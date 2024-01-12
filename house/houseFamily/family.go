package houseFamily

type Family struct {
	Member string
	Name   string
}

func AddMember() []Family {
	return []Family{
		{Member: "Mom", Name: "Vera"},
		{Member: "Dad", Name: "Ivan"},
		{Member: "Son", Name: "Max"},
		{Member: "Son", Name: "Alex"},
		{Member: "Baby", Name: "no name"},
	}
}
