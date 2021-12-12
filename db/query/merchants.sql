  -- name: CreateMerchant :one
INSERT INTO merchants (
  name, lat, long, logo
) VALUES(
  $1, $2, $3, $4
) 
RETURNING *;

  -- name: UpdateMerchant :one
UPDATE merchants
SET name = $2, long = $3, lat = $4
WHERE id = $1
RETURNING *;

  -- name: GetOneMerchantById :one
SELECT * from merchants 
WHERE id = $1 LIMIT 1;

  -- name: GetAllMerchants :many
SELECT * from merchants ORDER BY id ASC;

  -- name: SearchCertainMerchants :many
SELECT * from merchants WHERE name LIKE $1 ORDER BY id ASC;

  -- name: DeleteMerchant :exec
DELETE from merchants
WHERE id = $1;

  -- name: GetNearestMerchants :many
select * from (
SELECT  *,( 3959 * acos( cos( radians($1) ) * cos( radians( lat ) ) * cos( radians( long ) - radians($2) ) + sin( radians($1) ) * sin( radians( lat ) ) ) ) AS distance 
FROM merchants
) merchants
where distance < 5
ORDER BY distance
LIMIT 20;













