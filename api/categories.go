package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	db "github.com/AirellJordan98/hacktiv_ecommerce/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) createCategory(ctx *gin.Context) {
	var req createCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Bad Request", errors.New("category name required")))
		return
	}

	category, err := server.store.CreateCategory(context.Background(), req.Name)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("Internal Server Error", errors.New("something went wrong")))
		return
	}

	ctx.JSON(http.StatusCreated, category)

}

type updateCategoryRequest struct {
	Name string `json:"name"`
}

func (server *Server) updateCategory(ctx *gin.Context) {
	var req updateCategoryRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Bad Request", err))
		return
	}

	categoryId, err := strconv.Atoi(ctx.Param("categoryId"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Bad Request", errors.New("invalid params id")))
		return
	}

	arg := db.UpdateCategoryParams{
		ID:   int64(categoryId),
		Name: req.Name,
	}

	category, err := server.store.UpdateCategory(context.Background(), arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("Internal Server Error", err))
		return
	}

	ctx.JSON(http.StatusOK, category)

}

type categoryResponse struct {
	ID        int64        `json:"id"`
	Name      string       `json:"name"`
	Products  []db.Product `json:"Products"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

//GetOneCategory
func (server *Server) getOneCategoryById(ctx *gin.Context) {
	categoryId, err := strconv.Atoi(ctx.Param("categoryId"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Bad Request", errors.New("invalid params id")))
		return
	}

	category, err := server.store.GetOneCategoryById(context.Background(), int64(categoryId))

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse("Data Not Found", errors.New("category doens't exist")))
		return
	}

	productDatas := []db.Product{}

	sqlStmnt := "SELECT id, name, price, category_id, merchant_id, image, stock, created_at, updated_at from products WHERE category_id = $1"

	rows, err := queryDB.Query(sqlStmnt, int64(categoryId))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Internal Server Error", errors.New("something went wrong")))
		return
	}

	for rows.Next() {
		productData := db.Product{}
		err = rows.Scan(
			&productData.ID,
			&productData.Name,
			&productData.Price,
			&productData.CategoryID,
			&productData.MerchantID,
			&productData.Image,
			&productData.Stock,
			&productData.CreatedAt,
			&productData.UpdatedAt,
		)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse("Internal Server Error", errors.New("something went wrong")))
			return
		}

		productDatas = append(productDatas, productData)
	}

	categoryResponseData := categoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		Products:  productDatas,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, categoryResponseData)
}

func (server *Server) getAllCategories(ctx *gin.Context) {
	categoryResponseData := []categoryResponse{}
	categories, err := server.store.GetAllCategories(context.Background())

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Internal Server Error", errors.New("something went wrong")))
		return
	}

	productDatas := []db.Product{}
	sqlStmnt := "SELECT id, name, price, category_id, merchant_id, image, stock, created_at, updated_at from products WHERE category_id = $1"

	if len(categories) > 0 {
		for _, c := range categories {
			rows, err := queryDB.Query(sqlStmnt, c.ID)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, errorResponse("Internal Server Error", errors.New("something went wrong")))
				return
			}

			for rows.Next() {
				productData := db.Product{}
				err = rows.Scan(
					&productData.ID,
					&productData.Name,
					&productData.Price,
					&productData.CategoryID,
					&productData.MerchantID,
					&productData.Image,
					&productData.Stock,
					&productData.CreatedAt,
					&productData.UpdatedAt,
				)
				if err != nil {
					ctx.JSON(http.StatusBadRequest, errorResponse("Internal Server Error", errors.New("something went wrong")))
					return
				}

				productDatas = append(productDatas, productData)
			}
			categoryResponseData = append(categoryResponseData, categoryResponse{
				ID:        c.ID,
				Name:      c.Name,
				Products:  productDatas,
				CreatedAt: c.CreatedAt,
				UpdatedAt: c.UpdatedAt,
			})
		}
	}

	ctx.JSON(http.StatusOK, categoryResponseData)
}
