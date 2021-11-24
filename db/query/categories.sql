  -- name: CreateCategory :one
INSERT INTO categories (
  name
) VALUES(
  $1
) 
RETURNING *;

  -- name: UpdateCategory :one
UPDATE categories
SET name = $2
WHERE id = $1
RETURNING *;

  -- name: GetOneCategoryById :one
SELECT * from categories
WHERE id = $1 LIMIT 1;


  -- name: GetAllCategories :many
SELECT * from categories;


  -- name: DeleteCategory :exec
DELETE from categories
WHERE id = $1;


