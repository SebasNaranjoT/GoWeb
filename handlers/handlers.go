package handlers

import (
	"net/http"
	"strconv"

	"github.com/SebasNaranjoT/GoWeb.git/models"
	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	ctx.String(200, "pong")
}

func Products(ctx *gin.Context) {
	productos := models.Productos
	ctx.JSON(http.StatusOK, gin.H{
		"products": productos,
	})
}

func GetProductById(ctx *gin.Context) {
	productId, ok := strconv.Atoi(ctx.Param("id"))

	if ok != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "fail to parse id",
			"data":    nil,
		})
		return
	}

	var targetProduct models.Producto
	for _, producto := range models.Productos {
		if producto.Id == productId {
			targetProduct = producto
			break
		}
	}

	if targetProduct.Id != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Product succesfully founded",
			"product": targetProduct,
		})
		return
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Id not found",
			"product": nil,
		})
		return
	}
}

func ProdcutsGreaterThanPriceGt(ctx *gin.Context) {
	price, err := strconv.ParseFloat(ctx.Query("price"), 64)

	if err != nil {
		ctx.String(http.StatusBadRequest, "Something gone wrong")
		return
	}
	if price != 0 {
		var selectedProducts []models.Producto
		for _, product := range models.Productos {
			if product.Price > price {
				selectedProducts = append(selectedProducts, product)
			}
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message":  "Products filtered",
			"products": selectedProducts,
		})
	}
}
