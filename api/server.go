package api

import (
	"database/sql"

	db "github.com/AirellJordan98/hacktiv_ecommerce/db/sqlc"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

var queryDB *sql.DB

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store, q *sql.DB) *Server {
	queryDB = q
	server := &Server{store: store}
	router := gin.Default()

	router.Use(cors.Default())

	users := router.Group("/users")
	{
		users.POST("/register", server.userRegister)
		users.POST("/login", server.userLogin)
	}

	merchantRouter := router.Group("/merchants")
	{
		merchantRouter.POST("/", server.createMerchant)
		merchantRouter.GET("/", server.getAllMerchants)
		merchantRouter.GET("/:merchantId", server.getOneMerchantById)
		merchantRouter.PUT("/:merchantId", server.updateMerchant)
		merchantRouter.POST("/get-nearest-merchants", server.getNearestMerchants)
		merchantRouter.GET("/search-merchants", server.searchCertainMerchants)
	}

	categoryRouter := router.Group("/categories")
	{
		categoryRouter.GET("/", server.getAllCategories)
		categoryRouter.POST("/", server.createCategory)
		categoryRouter.GET("/:categoryId", server.getOneCategoryById)
	}

	productRouter := router.Group("/products")
	{
		productRouter.POST("/", server.createProduct)
		productRouter.GET("/", server.GetProducts)
		productRouter.GET("/:productId", server.GetOneProductById)
		productRouter.PUT("/:productId", server.updateProduct)
		productRouter.PATCH("/:productId", server.updateProductStock)
		productRouter.DELETE("/:productId", server.DeleteProductById)
	}

	cartRouter := router.Group("/carts")
	{
		cartRouter.POST("/", authentication(), server.createCart)
		cartRouter.PATCH("/:cartId", authentication(), server.CartUserAuthorization(), server.updateCart)
		cartRouter.DELETE("/:cartId", authentication(), server.CartUserAuthorization(), server.deleteCart)

		cartRouter.GET("/my-carts", authentication(), server.getCartsByUserId)

	}

	server.router = router

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(errCode string, errMsg error) gin.H {
	return gin.H{
		"errCode": errCode,
		"errMsg":  errMsg.Error(),
	}
}
