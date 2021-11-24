package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	db "github.com/AirellJordan98/hacktiv_ecommerce/db/sqlc"
	"github.com/AirellJordan98/hacktiv_ecommerce/util"
	"github.com/gin-gonic/gin"
)

type userRegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (server *Server) userRegister(ctx *gin.Context) {
	var req userRegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Bad Request", err))
		return
	}

	if len(req.Password) < int(8) {
		ctx.JSON(http.StatusBadRequest, errorResponse("Bad Request", errors.New("password has to have more than or equal to 8 characters")))
		return
	}

	hashPassword, err := util.HashPassword(req.Password)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse("Internal Server Error", errors.New("something went wrong")))
		return
	}

	arg := db.CreateUserParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashPassword,
	}

	user, err := server.store.CreateUser(context.Background(), arg)

	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, errorResponse("Internal Server Error", errors.New("something went wrong")))
		return
	}

	ctx.JSON(http.StatusCreated, user)

}

type userLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (server *Server) userLogin(ctx *gin.Context) {
	fmt.Println("Masuk user login")
	var req userLoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Bad Request", err))
		return
	}

	email := req.Email

	user, err := server.store.GetOneUserByEmail(context.Background(), email)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("Bad Request", errors.New("invalid email/password")))
		return
	}

	if valid := util.CheckPasswordHash(req.Password, user.Password); !valid {
		ctx.JSON(http.StatusBadRequest, errorResponse("Bad Request", errors.New("invalid email/password")))
		return
	}

	token := util.GenerateToken(uint(user.ID), user.Email)

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
