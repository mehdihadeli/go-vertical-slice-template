package dtos

type CreateProductRequestDto struct {
	Name        string  `json:"name"        validate:"required,min=3"`
	Description string  `json:"description" validate:"required,min=3"`
	Price       float64 `json:"price"       validate:"required,gt=0"`
}
