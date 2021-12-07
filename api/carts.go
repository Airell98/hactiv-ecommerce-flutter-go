package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	db "github.com/AirellJordan98/hacktiv_ecommerce/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type createCartRequest struct {
	MerchantID int32 `json:"merchant_id"`
	ProductID  int32 `json:"product_id"`
	Qty        int32 `json:"qty"`
}

func (server *Server) createCart(ctx *gin.Context) {
	var req createCartRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Bad Request", err))
		return
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := int32(userData["id"].(float64))

	product, err := server.store.GetOneProductById(context.Background(), int64(req.ProductID))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("Internal Server Error", errors.New("something went wrong")))
		return
	}

	if req.Qty <= 0 {
		ctx.JSON(http.StatusInternalServerError, errorResponse("Bad Request", errors.New("quantity has to be more than or equal to one")))
		return
	}

	if product.Stock < req.Qty {
		ctx.JSON(http.StatusInternalServerError, errorResponse("Bad Request", fmt.Errorf("insufficient product stock, only %d left", product.Stock)))
		return
	}

	totalPrice := req.Qty * product.Price
	_ = totalPrice

	findCart, err := server.store.GetOneCartByUserIdAndProductId(context.Background(), db.GetOneCartByUserIdAndProductIdParams{
		UserID:    userID,
		ProductID: req.ProductID,
	})
	_ = findCart

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			arg := db.CreateCartParams{
				UserID:     userID,
				MerchantID: req.MerchantID,
				ProductID:  req.ProductID,
				Qty:        req.Qty,
				TotalPrice: totalPrice,
			}

			cart, err := server.store.CreateCart(context.Background(), arg)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse("Internal Server Error", err))
				return
			}

			ctx.JSON(http.StatusCreated, cart)
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse("Internal Server Error", err))
		return
	}
	arg := db.UpdateCartQtyParams{
		ID:         findCart.ID,
		Qty:        req.Qty,
		TotalPrice: totalPrice,
	}

	updatedCart, err := server.store.UpdateCartQty(context.Background(), arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("Internal Server Error", errors.New("something went wrong")))
		return
	}

	ctx.JSON(http.StatusOK, updatedCart)

}

type updateCartRequest struct {
	Qty       int32 `json:"qty" binding:"required"`
	ProductID int32 `json:"product_id" binding:"required"`
}

func (server *Server) updateCart(ctx *gin.Context) {

	var req updateCartRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Bad Request", err))
		return
	}

	cartId, err := strconv.Atoi(ctx.Param("cartId"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Bad Request", errors.New("invalid params id")))
		return
	}

	if req.Qty == 0 {
		ctx.JSON(http.StatusBadRequest, errorResponse("Bad Request", errors.New("at least you're purchasing one product")))
		return
	}
	// fmt.Println("product id", req.ProductID)

	product, err := server.store.GetOneProductById(context.Background(), int64(req.ProductID))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Internal Server Error", err))
		return
	}

	if product.Stock < req.Qty {
		ctx.JSON(http.StatusInternalServerError, errorResponse("Bad Request", fmt.Errorf("insufficient product stock, only %d left", product.Stock)))
		return
	}

	totalPrice := product.Price * req.Qty

	arg := db.UpdateCartQtyParams{
		ID:         int64(cartId),
		Qty:        req.Qty,
		TotalPrice: totalPrice,
	}

	cart, err := server.store.UpdateCartQty(context.Background(), arg)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Internal Server Error", err))
		return
	}

	ctx.JSON(http.StatusOK, cart)
}

func (server *Server) getCartsByUserId(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := int32(userData["id"].(float64))

	carts, err := server.store.GetCartsByUserId(context.Background(), userID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Internal Server Error", err))
		return
	}

	cartResponseData := []interface{}{}

	for _, c := range carts {
		d := map[string]interface{}{
			"id":            c.ID,
			"product_id":    c.ProductID,
			"product_image": "",
			"qty":           c.Qty,
			"total_price":   c.TotalPrice,
			"user_id":       c.UserID,
			"created_at":    c.CreatedAt,
			"updated_at":    c.UpdatedAt,
			"merchant_name": "",
			"product_name":  "",
			"product_price": 0,
			"product_stock": 0,
		}

		if c.MerchantName.Valid {
			d["merchant_name"] = c.MerchantName.String
		}

		if c.ProductName.Valid {
			d["product_name"] = c.ProductName.String
		}

		if c.ProductPrice.Valid {
			d["product_price"] = c.ProductPrice.Int32
		}

		if c.ProductStock.Valid {
			d["product_stock"] = c.ProductStock.Int32
		}

		if c.ProductImage.Valid {
			d["product_image"] = c.ProductImage.String
		}

		cartResponseData = append(cartResponseData, d)

	}

	ctx.JSON(http.StatusOK, cartResponseData)
}

func (server *Server) deleteCart(ctx *gin.Context) {
	cartId, err := strconv.Atoi(ctx.Param("cartId"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Bad Request", errors.New("invalid params id")))
		return
	}

	err = server.store.DeleteCart(context.Background(), int64(cartId))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Internal Server Error", err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("cart with id %d has been successfully deleted", cartId),
	})
}
