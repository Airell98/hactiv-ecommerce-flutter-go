// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
)

type Querier interface {
	CreateCart(ctx context.Context, arg CreateCartParams) (Cart, error)
	CreateCategory(ctx context.Context, name string) (Category, error)
	CreateMerchant(ctx context.Context, arg CreateMerchantParams) (Merchant, error)
	CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error)
	DeleteCart(ctx context.Context, id int64) error
	DeleteCategory(ctx context.Context, id int64) error
	DeleteMerchant(ctx context.Context, id int64) error
	DeleteProduct(ctx context.Context, id int64) error
	DeleteUser(ctx context.Context, id int64) error
	GetAllCategories(ctx context.Context) ([]Category, error)
	GetAllMerchants(ctx context.Context) ([]Merchant, error)
	GetAllProducts(ctx context.Context) ([]Product, error)
	GetCartsByUserId(ctx context.Context, userID int32) ([]GetCartsByUserIdRow, error)
	GetNearestMerchants(ctx context.Context, arg GetNearestMerchantsParams) ([]GetNearestMerchantsRow, error)
	GetOneCartById(ctx context.Context, id int64) (GetOneCartByIdRow, error)
	GetOneCartByUserId(ctx context.Context, userID int32) (Cart, error)
	GetOneCartByUserIdAndProductId(ctx context.Context, arg GetOneCartByUserIdAndProductIdParams) (Cart, error)
	GetOneCategoryById(ctx context.Context, id int64) (Category, error)
	GetOneMerchantById(ctx context.Context, id int64) (Merchant, error)
	GetOneProductById(ctx context.Context, id int64) (Product, error)
	GetOneUserByEmail(ctx context.Context, email string) (User, error)
	UpdateCartQty(ctx context.Context, arg UpdateCartQtyParams) (Cart, error)
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error)
	UpdateMerchant(ctx context.Context, arg UpdateMerchantParams) (Merchant, error)
	UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error)
	UpdateProductStock(ctx context.Context, arg UpdateProductStockParams) (Product, error)
	UpdateUser(ctx context.Context, email string) (User, error)
}

var _ Querier = (*Queries)(nil)
