  -- name: CreateCart :one
INSERT INTO carts (
  user_id, merchant_id, product_id, qty, total_price
) VALUES(
  $1, $2, $3, $4, $5
) 
RETURNING *;


  -- name: GetCartsByUserId :many
SELECT c.id, c.user_id, c.product_id, c.qty, c.total_price, m.name as merchant_name, p.name as product_name, p.image as product_image, p.price as product_price, p.stock as product_stock, c.created_at, c.updated_at  from carts as c 
LEFT JOIN merchants as m on m.id = c.merchant_id
LEFT JOIN products as p on p.id = c.product_id
WHERE c.user_id = $1;


  -- name: GetOneCartByUserId :one
SELECT * from carts
WHERE user_id = $1;

  -- name: UpdateCartQty :one
UPDATE carts
SET qty = $2, total_price = $3
WHERE id = $1
RETURNING *;


  -- name: DeleteCart :exec
DELETE from carts
WHERE id = $1;
