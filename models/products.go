package models

//import "github.com/SebasNaranjoT/GoWeb.git/handlers"

type Producto struct {
	Id          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
	CodeValue   string  `json:"code_value" validate:"required"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}

var Productos []Producto

func NewProduct(id int, name string, quantity int, code string, published bool, exp string, price float64) {
	producto := Producto{
		id,
		name,
		quantity,
		code,
		published,
		exp,
		price,
	}

	Productos = append(Productos, producto)
}
