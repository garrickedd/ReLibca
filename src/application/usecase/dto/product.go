package dto

type Product struct {
	Id          int
	Name        string
	Description string
	Status      bool
	TypeId      int
	PropertyId  int
}

type CreateProduct struct {
	Name        string
	Description string
	Status      bool
	TypeId      int
	PropertyId  int
}

type UpdateProduct struct {
	Name        string
	Description string
	Status      bool
	TypeId      int
	PropertyId  int
}
