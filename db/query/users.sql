  -- name: CreateUser :one
INSERT INTO users (
  name, 
  email, 
  password
) VALUES(
  $1, $2, $3
) 
RETURNING id, name, email, created_at, updated_at;

  -- name: GetOneUserByEmail :one
SELECT * from users
WHERE email = $1 LIMIT 1;

  -- name: UpdateUser :one
SELECT * from users
WHERE email = $1 LIMIT 1;

  -- name: DeleteUser :exec
DELETE from users
WHERE id = $1;