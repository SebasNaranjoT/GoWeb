package models

type Producto struct {
	Id          int
	Name        string
	Quantity    float64
	CodeValue   string
	IsPublished bool
	Expiration  string
	Price       float64
}

var Productos []Producto
