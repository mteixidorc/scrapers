package core

type Room struct {
	Name        string
	Description string
	Price       float32
}

func NewRoom(name, description string, price float32) *Room {
	return &Room{
		Name:        name,
		Description: description,
		Price:       price,
	}
}

type ByPrice []*Room

func (a ByPrice) Len() int {
	return len(a)
}

func (a ByPrice) Less(i, j int) bool {
	return a[i].Price < a[j].Price
}

func (a ByPrice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
