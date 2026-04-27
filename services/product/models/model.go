package models

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type ProductFilter struct {
	Id *int
}

type ProductInput struct {
	Name        string  `json:"name" validate:"required"`
	Price       float64 `json:"price" validate:"required,number"`
	Description string  `json:"description" validate:"required"`
}

type ProductUpdateInput struct {
	Name        *string  `json:"name"`
	Price       *float64 `json:"price" validate:"omitempty,number"`
	Description *string  `json:"description"`
}

type ProductUpdatePutInput struct {
	Name        string  `json:"name" validate:"required"`
	Price       float64 `json:"price" validate:"required,number"`
	Description string  `json:"description" validate:"required"`
}
