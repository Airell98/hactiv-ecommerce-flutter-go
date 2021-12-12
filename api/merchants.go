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

type createMerchantRequest struct {
	Name string `json:"name" binding:"required"`
	Lat  string `json:"lat" binding:"required"`
	Long string `json:"long" binding:"required"`
	Logo string `json:"logo" binding:"required"`
}

func (server *Server) createMerchant(ctx *gin.Context) {
	var req createMerchantRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Bad Request", err))
		return
	}

	arg := db.CreateMerchantParams{
		Name: req.Name,
		Lat:  req.Lat,
		Long: req.Long,
		Logo: req.Logo,
	}

	merchant, err := server.store.CreateMerchant(context.Background(), arg)

	if err != nil {
		fmt.Println("Error: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, errorResponse("Internal Server Error", errors.New("something went wrong")))
		return
	}

	ctx.JSON(http.StatusCreated, merchant)

}

type updateMerchantRequest struct {
	Name string `json:"name" binding:"required"`
	Lat  string `json:"lat" binding:"required"`
	Long string `json:"long" binding:"required"`
	Logo string `json:"logo" binding:"required"`
}

func (server *Server) updateMerchant(ctx *gin.Context) {
	var req updateMerchantRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Bad Request", err))
		return
	}
	merchantId, err := strconv.Atoi(ctx.Param("merchantId"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Bad Request", errors.New("invalid params id")))
		return
	}

	arg := db.UpdateMerchantParams{
		ID:   int64(merchantId),
		Name: req.Name,
		Long: req.Long,
		Lat:  req.Lat,
	}

	merchant, err := server.store.UpdateMerchant(context.Background(), arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("Internal Server Error", errors.New("something went wrong")))
		return
	}

	ctx.JSON(http.StatusOK, merchant)
}

type MerchantResponse struct {
	ID        int64        `json:"id"`
	Name      string       `json:"name"`
	Lat       string       `json:"lat"`
	Long      string       `json:"long"`
	Logo      string       `json:"logo"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	Products  []db.Product `json:"products"`
}

func (server *Server) getOneMerchantById(ctx *gin.Context) {
	merchantId, err := strconv.Atoi(ctx.Param("merchantId"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Bad Request", errors.New("invalid params id")))
		return
	}

	merchant, err := server.store.GetOneMerchantById(context.Background(), int64(merchantId))

	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse("Data Not Found", errors.New("merchant doesn't exist")))
		return
	}
	productDatas := []db.Product{}

	sqlStmnt := "SELECT id, name, price, category_id, merchant_id, image, stock, created_at, updated_at from products WHERE merchant_id = $1"

	rows, err := queryDB.Query(sqlStmnt, int64(merchantId))

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

	res := MerchantResponse{
		ID:        merchant.ID,
		Name:      merchant.Name,
		Long:      merchant.Long,
		Lat:       merchant.Lat,
		Logo:      merchant.Logo,
		CreatedAt: merchant.CreatedAt,
		UpdatedAt: merchant.UpdatedAt,
		Products:  productDatas,
	}
	ctx.JSON(http.StatusOK, res)
}

func (server *Server) getAllMerchants(ctx *gin.Context) {

	merchant, err := server.store.GetAllMerchants(context.Background())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("Internal Server Error", errors.New("something went wrong")))
		return
	}

	ctx.JSON(http.StatusOK, merchant)
}

type getNearestMerchantsRequest struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

func (server *Server) getNearestMerchants(ctx *gin.Context) {
	var req getNearestMerchantsRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Bad Request", err))
		return
	}

	arg := db.GetNearestMerchantsParams{
		Radians:   req.Lat,
		Radians_2: req.Long,
	}

	fmt.Println(arg)
	merchants, err := server.store.GetNearestMerchants(context.Background(), arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("Internal Server Error", err))
		return
	}

	ctx.JSON(http.StatusOK, merchants)

}

func (server *Server) searchCertainMerchants(ctx *gin.Context) {
	var merchantName string = ctx.Param("merchantId")

	merchantName = "%" + merchantName + "%"
	merchants, err := server.store.SearchCertainMerchants(context.Background(), merchantName)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Bad Request", err))
		return
	}

	ctx.JSON(http.StatusOK, merchants)
}
