package models

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type ProductFilter struct {
	Id int `json:'id' validate:"required"`
}

type CreateProductInput struct {
	Name        string  `json:"name" validate:"required"`
	Price       float64 `json:"price" validate:"required,number"`
	Description string  `json:"description" validate:"required"`
}

type PatchProductInput struct {
	Name        *string  `json:"name"`
	Price       *float64 `json:"price" validate:"omitempty,number"`
	Description *string  `json:"description"`
}

type PatchProductInputGraphql struct {
	Id          int      `json:"id" validate:"required"`
	Name        *string  `json:"name"`
	Price       *float64 `json:"price" validate:"omitempty,number"`
	Description *string  `json:"description"`
}

type PutProductInput struct {
	Name        string  `json:"name" validate:"required"`
	Price       float64 `json:"price" validate:"required,number"`
	Description string  `json:"description" validate:"required"`
}

type DeleteProductInput struct {
	Id int `json:"id" validate:"required"`
}
