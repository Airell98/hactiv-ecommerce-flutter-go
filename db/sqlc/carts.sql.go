// Code generated by sqlc. DO NOT EDIT.
// source: carts.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createCart = `-- name: CreateCart :one
INSERT INTO carts (
  user_id, merchant_id, product_id, qty, total_price
) VALUES(
  $1, $2, $3, $4, $5
) 
RETURNING id, user_id, merchant_id, product_id, qty, total_price, created_at, updated_at
`

type CreateCartParams struct {
	UserID     int32 `json:"user_id"`
	MerchantID int32 `json:"merchant_id"`
	ProductID  int32 `json:"product_id"`
	Qty        int32 `json:"qty"`
	TotalPrice int32 `json:"total_price"`
}

func (q *Queries) CreateCart(ctx context.Context, arg CreateCartParams) (Cart, error) {
	row := q.db.QueryRowContext(ctx, createCart,
		arg.UserID,
		arg.MerchantID,
		arg.ProductID,
		arg.Qty,
		arg.TotalPrice,
	)
	var i Cart
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.MerchantID,
		&i.ProductID,
		&i.Qty,
		&i.TotalPrice,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteCart = `-- name: DeleteCart :exec
DELETE from carts
WHERE id = $1
`

func (q *Queries) DeleteCart(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCart, id)
	return err
}

const getCartsByUserId = `-- name: GetCartsByUserId :many
SELECT c.id, c.user_id, c.product_id, c.qty, c.total_price, m.name as merchant_name, p.name as product_name, p.image as product_image, p.price as product_price, p.stock as product_stock, c.created_at, c.updated_at  from carts as c 
LEFT JOIN merchants as m on m.id = c.merchant_id
LEFT JOIN products as p on p.id = c.product_id
WHERE c.user_id = $1
`

type GetCartsByUserIdRow struct {
	ID           int64          `json:"id"`
	UserID       int32          `json:"user_id"`
	ProductID    int32          `json:"product_id"`
	Qty          int32          `json:"qty"`
	TotalPrice   int32          `json:"total_price"`
	MerchantName sql.NullString `json:"merchant_name"`
	ProductName  sql.NullString `json:"product_name"`
	ProductImage sql.NullString `json:"product_image"`
	ProductPrice sql.NullInt32  `json:"product_price"`
	ProductStock sql.NullInt32  `json:"product_stock"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

func (q *Queries) GetCartsByUserId(ctx context.Context, userID int32) ([]GetCartsByUserIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getCartsByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetCartsByUserIdRow{}
	for rows.Next() {
		var i GetCartsByUserIdRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ProductID,
			&i.Qty,
			&i.TotalPrice,
			&i.MerchantName,
			&i.ProductName,
			&i.ProductImage,
			&i.ProductPrice,
			&i.ProductStock,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOneCartByUserId = `-- name: GetOneCartByUserId :one
SELECT id, user_id, merchant_id, product_id, qty, total_price, created_at, updated_at from carts
WHERE user_id = $1
`

func (q *Queries) GetOneCartByUserId(ctx context.Context, userID int32) (Cart, error) {
	row := q.db.QueryRowContext(ctx, getOneCartByUserId, userID)
	var i Cart
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.MerchantID,
		&i.ProductID,
		&i.Qty,
		&i.TotalPrice,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getOneCartByUserIdAndProductId = `-- name: GetOneCartByUserIdAndProductId :one
SELECT id, user_id, merchant_id, product_id, qty, total_price, created_at, updated_at from carts
WHERE user_id = $1 AND product_id = $2
`

type GetOneCartByUserIdAndProductIdParams struct {
	UserID    int32 `json:"user_id"`
	ProductID int32 `json:"product_id"`
}

func (q *Queries) GetOneCartByUserIdAndProductId(ctx context.Context, arg GetOneCartByUserIdAndProductIdParams) (Cart, error) {
	row := q.db.QueryRowContext(ctx, getOneCartByUserIdAndProductId, arg.UserID, arg.ProductID)
	var i Cart
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.MerchantID,
		&i.ProductID,
		&i.Qty,
		&i.TotalPrice,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateCartQty = `-- name: UpdateCartQty :one
UPDATE carts
SET qty = $2, total_price = $3
WHERE id = $1
RETURNING id, user_id, merchant_id, product_id, qty, total_price, created_at, updated_at
`

type UpdateCartQtyParams struct {
	ID         int64 `json:"id"`
	Qty        int32 `json:"qty"`
	TotalPrice int32 `json:"total_price"`
}

func (q *Queries) UpdateCartQty(ctx context.Context, arg UpdateCartQtyParams) (Cart, error) {
	row := q.db.QueryRowContext(ctx, updateCartQty, arg.ID, arg.Qty, arg.TotalPrice)
	var i Cart
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.MerchantID,
		&i.ProductID,
		&i.Qty,
		&i.TotalPrice,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
