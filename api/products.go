package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	db "github.com/AirellJordan98/hacktiv_ecommerce/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createProductRequest struct {
	Name       string `json:"name" binding:"required"`
	Price      int32  `json:"price"  binding:"required"`
	CategoryID int32  `json:"category_id" binding:"required"`
	MerchantID int32  `json:"merchant_id" binding:"required"`
	Image      string `json:"image" binding:"required"`
}

func (server *Server) createProduct(ctx *gin.Context) {
	var req createProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Bad Request", err))
		return
	}

	arg := db.CreateProductParams{
		Name:       req.Name,
		Price:      req.Price,
		CategoryID: req.CategoryID,
		MerchantID: req.MerchantID,
		Image:      req.Image,
		Stock:      100,
	}

	product, err := server.store.CreateProduct(context.Background(), arg)

	if err != nil {
		fmt.Println("Error: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, errorResponse("Internal Server Error", errors.New("something went wrong")))
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (server *Server) GetOneProductById(ctx *gin.Context) {
	productId, err := strconv.Atoi(ctx.Param("productId"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Bad Request", errors.New("invalid params id")))
		return
	}

	product, err := server.store.GetOneProductById(context.Background(), int64(productId))

	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse("Data Not Found", errors.New("product doesn't exist")))
		return
	}

	ctx.JSON(http.StatusOK, product)

}

func (server *Server) GetProducts(ctx *gin.Context) {
	products, err := server.store.GetAllProducts(context.Background())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("Internal Server Error", errors.New("something went wrong")))
		return
	}

	ctx.JSON(http.StatusOK, products)

}

type updateProductRequest struct {
	Name  string `json:"name"`
	Price int32  `json:"price"`
	Image string `json:"image"`
}

func (server *Server) updateProduct(ctx *gin.Context) {
	var req updateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Bad Request", err))
		return
	}

	productId, err := strconv.Atoi(ctx.Param("productId"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Bad Request", errors.New("invalid params id")))
		return
	}

	arg := db.UpdateProductParams{
		ID:    int64(productId),
		Name:  req.Name,
		Price: req.Price,
		Image: req.Image,
	}

	product, err := server.store.UpdateProduct(context.Background(), arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("Internal Server Error", errors.New("something went wrong")))
		return
	}

	ctx.JSON(http.StatusOK, product)
}

type updateProductStockRequest struct {
	Stock int32 `json:"stock" binding:"required"`
}

func (server *Server) updateProductStock(ctx *gin.Context) {
	var req updateProductStockRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Bad Request", err))
		return
	}

	if req.Stock < 0 {
		ctx.JSON(http.StatusBadRequest, errorResponse("Bad Request", errors.New("stock cannot be less than 0")))
		return
	}

	productId, err := strconv.Atoi(ctx.Param("productId"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Bad Request", errors.New("invalid params id")))
		return
	}

	arg := db.UpdateProductStockParams{
		ID:    int64(productId),
		Stock: req.Stock,
	}

	product, err := server.store.UpdateProductStock(context.Background(), arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("Internal Server Error", errors.New("something went wrong")))
		return
	}

	_ = product

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Stock has been successfully updated",
	})
}

func (server *Server) DeleteProductById(ctx *gin.Context) {
	productId, err := strconv.Atoi(ctx.Param("productId"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Bad Request", errors.New("invalid params id")))
		return
	}

	err = server.store.DeleteProduct(context.Background(), int64(productId))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("Internal Server Error", errors.New("product doesn't exist")))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Product with id %d has been successfully deleted", productId),
	})

}
