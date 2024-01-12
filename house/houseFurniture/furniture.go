package houseFurniture

type Furniture struct {
	Name string
	Size float64
}

func AddFurniture() []Furniture {
	return []Furniture{
		{Name: "Bed", Size: 3.48},
		{Name: "Table", Size: 8},
		{Name: "Carpet", Size: 10.456},
		{Name: "Window", Size: 6},
		{Name: "TV", Size: 5.6},
	}
}
