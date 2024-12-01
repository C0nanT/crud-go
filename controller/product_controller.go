package controller

import (
	"go-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) ProductController {
	return ProductController{
		ProductUsecase: usecase,
	}
}

func (p *ProductController) GetProducts(ctx *gin.Context) {
	//GetProducts
	products, err := p.ProductUsecase.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, products)
}