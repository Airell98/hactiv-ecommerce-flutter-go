  -- name: CreateProduct :one
INSERT INTO products (
  name, price, category_id, merchant_id, image, stock
) VALUES(
  $1, $2, $3, $4, $5, $6
) 
RETURNING *;


  -- name: UpdateProduct :one
UPDATE products
SET name = $2, price = $3, image = $4
WHERE id = $1
RETURNING *;


  -- name: UpdateProductStock :one
UPDATE products
SET stock = $2
WHERE id = $1
RETURNING *;

  -- name: GetOneProductById :one
SELECT * from products
WHERE id = $1 LIMIT 1;

  -- name: GetAllProducts :many
SELECT * from products;


  -- name: DeleteProduct :exec
DELETE from products
WHERE id = $1;


