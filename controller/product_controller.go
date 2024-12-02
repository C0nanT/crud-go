package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

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
	products, err := p.ProductUsecase.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *ProductController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	newProduct, err := p.ProductUsecase.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, newProduct)
}

func (p *ProductController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("id")

	if (id == "") {
		response := model.Response{
			Message: "ID is required",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{
			Message: "ID must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.ProductUsecase.GetProductById(productId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response{
			Message: "Product not found",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *ProductController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	if (id == "") {
		response := model.Response{
			Message: "ID is required",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{
			Message: "ID must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = p.ProductUsecase.DeleteProduct(productId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	response := model.Response{
		Message: "Product deleted",
	}

	ctx.JSON(http.StatusOK, response)
}

func (p *ProductController) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	if (id == "") {
		response := model.Response{
			Message: "ID is required",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{
			Message: "ID must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var product model.Product
	err = ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	updatedProduct, err := p.ProductUsecase.UpdateProduct(productId, product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, updatedProduct)
}