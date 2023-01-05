package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/SebasNaranjoT/GoWeb.git/handlers"
	"github.com/SebasNaranjoT/GoWeb.git/models"
	"github.com/gin-gonic/gin"
)

func readJson() {
	data, err := ioutil.ReadFile("products.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(data, &models.Productos)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	readJson()
	server := gin.Default()
	server.GET("/ping", handlers.Ping)
	server.GET("/products", handlers.Products)
	server.GET("/products/:id", handlers.GetProductById)
	server.GET("/products/search", handlers.ProdcutsGreaterThanPriceGt)
	server.POST("/products/new", handlers.CreateProduct)
	server.Run()
}
