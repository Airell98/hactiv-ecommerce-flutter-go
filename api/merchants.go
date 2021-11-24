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

	ctx.JSON(http.StatusOK, merchant)
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
		ctx.JSON(http.StatusBadRequest, errorResponse("Internal Server Error", err))
		return
	}

	ctx.JSON(http.StatusOK, merchants)

}
