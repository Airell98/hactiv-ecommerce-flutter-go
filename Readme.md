# RESTful endpoints
```
Heroku URL: https://hacktiv-ecommerce.herokuapp.com/
```

<!-- --- -->

<h1 style="color: orange; font-weight: 700;">Users</h1>

### POST /users/login

_Request Header_

```
Not needed
```

_Request Body_

```
{
  "email" : "string<your email>",
  "password" : "string<your password>"

}
```

_Response (200)_

```
{ 
  "token" : "<users's token>",
}
```

_Response (400)_

```
{ "error"  : "Bad Request"
  "message": "invalid Email / Password"
}
```

_Response (500)_

```
{  "error" : "Internal Server Error",
  "message": "Something went wrong"
}
```

---


### POST /users/register

_Request Header_

```
Not needed
```

_Request Body_

```
{
  "name" : "string<user's name>",
  "email" : "string<user's email (unique)>",
  "password" : "string<user's password>"

}
```

_Response (201)_

```
{ 
  "id": 1,
  "name": "john doe",
  "email": "johndoe@mail.com",
  "created_at": 2021-11-15 18:12:50.292563+07,
  "updated_at": 2021-11-15 18:12:50.292563+07
}
```

_Response (400)_

```
{ 
  "error"  : "Bad Request"
  "message": "name, email, password required"
}
```

_Response (400)_

```
{ 
  "error"  : "Bad Request"
  "message": "password has to have more than or equal to 8 characters"
}
```

_Response (400)_

```
{ 
  "error"  : "Bad Request"
  "message": "email has been taken"
}
```

_Response (500)_

```
{  "error" : "Internal Server Error",
  "message": "Something went wrong"
}
```

---

<h1 style="color: orange;">Categories</h1>

### POST /categories
**(create a new category)**

_Request Header_

```
Not needed
```

_Request Body_

```
{
  "name" : "string<category's name>"
}
```

_Response (201)_

```
{ 
  "id": 1,
  "name": "Beverages",
  "created_at": 2021-11-15 18:12:50.292563+07,
  "updated_at": 2021-11-15 18:12:50.292563+07
}
```

_Response (400)_

```
{ "error"  : "Bad Request"
  "message": "category name required"
}
```

_Response (500)_

```
{  "error" : "Internal Server Error",
  "message": "something went wrong"
}
```

---

### PATCH /categories/:categoryId
**(update category's name by categoryId)**

_Request Header_

```
Not needed
```

_Request Body_

```
{
  "name" : "string<category's name>"
}
```

_Response (200)_

```
{ 
  "id": 1,
  "name": "Beverages",
  "created_at": 2021-11-15 18:12:50.292563+07,
  "updated_at": 2021-11-15 18:12:50.292563+07
}
```

_Response (400)_

```
{ "error"  : "Bad Request"
  "message": "category name required"
}
```

_Response (500)_

```
{  "error" : "Internal Server Error",
  "message": "something went wrong"
}
```

---

### GET /categories/:categoryId
**(get one category data by category id)**

_Request Header_

```
Not needed
```

_Request Body_

```
Not needed
```

_Response (200)_

```
{
    "id": 1,
    "name": "Beverages",
    "Products": [
        {
            "id": 1,
            "name": "Ultra Milk",
            "price": 6000,
            "category_id": 1,
            "merchant_id": 1,
            "image": "http://static.bmdstatic.com/pk/product/large/5f87d8756ae79.jpg",
            "stock": 20,
            "created_at": "2021-11-15T14:45:27.93129+08:00",
            "updated_at": "2021-11-15T14:45:27.93129+08:00"
        }
    ],
    "created_at": "2021-11-15T14:32:41.221813+08:00",
    "updated_at": "2021-11-15T14:32:41.221813+08:00"
}
```

_Response (400)_

```
{ "error"  : "Bad Request"
  "message": "invalid id params"
}
```

_Response (404)_

```
{ "error"  : "Data Not Found"
  "message": "category doesn't exist"
}
```

_Response (500)_

```
{  "error" : "Internal Server Error",
  "message": "something went wrong"
}
```

---

### GET /categories
**(get all category datas)**

_Request Header_

```
Not needed
```

_Request Body_

```
Not needed
```

_Response (200)_

```
[
  {
    "id": 1,
    "name": "Beverages",
    "Products": [
        {
            "id": 1,
            "name": "Ultra Milk",
            "price": 6000,
            "category_id": 1,
            "merchant_id": 1,
            "image": "http://static.bmdstatic.com/pk/product/large/5f87d8756ae79.jpg",
            "stock": 20,
            "created_at": "2021-11-15T14:45:27.93129+08:00",
            "updated_at": "2021-11-15T14:45:27.93129+08:00"
        }
    ],
    "created_at": "2021-11-15T14:32:41.221813+08:00",
    "updated_at": "2021-11-15T14:32:41.221813+08:00"
  },
  {
    "id": 2,
    "name": "Junk Foods",
    "Products": [
        {
            "id": 1,
            "name": "Hamburgeria",
            "price": 20000,
            "category_id": 2,
            "merchant_id": 1,
            "image": "http://static.bmdstatic.com/pk/product/large/5f87d8756ae79.jpg",
            "stock": 40,
            "created_at": "2021-11-15T14:45:27.93129+08:00",
            "updated_at": "2021-11-15T14:45:27.93129+08:00"
        }
    ],
    "created_at": "2021-11-15T14:32:41.221813+08:00",
    "updated_at": "2021-11-15T14:32:41.221813+08:00"
  }
]
```

_Response (400)_

```
{ "error"  : "Bad Request"
  "message": "invalid id params"
}
```

_Response (404)_

```
{ "error"  : "Data Not Found"
  "message": "category doesn't exist"
}
```

_Response (500)_

```
{  "error" : "Internal Server Error",
  "message": "something went wrong"
}
```

---


<h1 style="color: orange; font-weight: 700;">Products</h1>

### POST /products
**(create a new product data)**

_Request Header_

```
Not needed
```

_Request Body_

```
{   
    "name": "string<product's name>",
    "price": integer<product's price>,
    "category_id": integer<product's category_id>,
    "merchant_id": integer<product's merchant_id>,
    "image": "string<product's image url>"
}
```

_Response (201)_

```
{
    "id": 1,
    "name": "Ultra Milk",
    "price": 6000,
    "category_id": 1,
    "merchant_id": 1,
    "image": "http://static.bmdstatic.com/pk/product/large/5f87d8756ae79.jpg",
    "stock": 20,
    "created_at": "2021-11-15T14:45:27.93129+08:00",
    "updated_at": "2021-11-15T14:45:27.93129+08:00"
}
```

_Response (400)_

```
{ "error"  : "Bad Request"
  "message": "product name, price, category_id, merchant_id, and image required"
}
```

_Response (500)_

```
{  "error" : "Internal Server Error",
  "message": "something went wrong"
}
```

---

### GET /products/:productId
**(get one product data by productId)**

_Request Header_

```
Not needed
```

_Request Body_

```
Not needed
```

_Response (200)_

```
{
    "id": 1,
    "name": "Ultra Milk",
    "price": 6000,
    "category_id": 1,
    "merchant_id": 1,
    "image": "http://static.bmdstatic.com/pk/product/large/5f87d8756ae79.jpg",
    "stock": 20,
    "created_at": "2021-11-15T14:45:27.93129+08:00",
    "updated_at": "2021-11-15T14:45:27.93129+08:00"
}
```

_Response (400)_

```
{ "error"  : "Bad Request"
  "message": "invalid id params"
}
```

_Response (404)_

```
{ "error"  : "Data Not Found"
  "message": "product doesn't exist"
}
```

_Response (500)_

```
{  "error" : "Internal Server Error",
  "message": "something went wrong"
}
```

---

### PUT /products/:productId
**(update one product data by productId)**

_Request Header_

```
Not needed
```

_Request Body_

```
{   
    "name": "string<product's name>",
    "price": integer<product's price>,
    "image": "string<product's image url>"
}
```

_Response (200)_

```
{
    "id": 2,
    "name": "Milo",
    "price": 4000,
    "category_id": 1,
    "merchant_id": 1,
    "image": "http://static.bmdstatic.com/pk/product/large/5f87d8756ae79.jpg",
    "stock": 80,
    "created_at": "2021-11-15T14:45:27.93129+08:00",
    "updated_at": "2021-11-15T14:45:27.93129+08:00"
}
```

_Response (400)_

```
{ "error"  : "Bad Request"
  "message": "invalid id params"
}
```

_Response (404)_

```
{ "error"  : "Data Not Found"
  "message": "product doesn't exist"
}
```

_Response (500)_

```
{  "error" : "Internal Server Error",
  "message": "something went wrong"
}
```

---

### PATCH /products/:productId
**(update product stock by productId)**

_Request Header_

```
Not needed
```

_Request Body_

```
{   
    "name": "string<product's name>",
    "price": integer<product's price>,
    "image": "string<product's image url>"
}
```

_Response (200)_

```
{
    "id": 2,
    "name": "Milo",
    "price": 4000,
    "category_id": 1,
    "merchant_id": 1,
    "image": "http://static.bmdstatic.com/pk/product/large/5f87d8756ae79.jpg",
    "stock": 80,
    "created_at": "2021-11-15T14:45:27.93129+08:00",
    "updated_at": "2021-11-15T14:45:27.93129+08:00"
}
```

_Response (400)_

```
{ "error"  : "Bad Request"
  "message": "invalid id params"
}
```

_Response (400)_

```
{ "error"  : "Bad Request"
  "message": "stock cannot be less than 0"
}
```


_Response (500)_

```
{  "error" : "Internal Server Error",
  "message": "something went wrong"
}
```

---

### DELETE /products/:productId
**(delete one product by productId)**

_Request Header_

```
Not needed
```

_Request Body_

```
Not needed
```

_Response (200)_

```
{
  "message": "Product with id 2 has been successfully deleted"
}
```

_Response (400)_

```
{  "error" : "Bad Request",
  "message": "invalid id params"
}
```

_Response (500)_

```
{  "error" : "Internal Server Error",
  "message": "something went wrong"
}
```

---

---

### GET /products
**(get all product datas)**

_Request Header_

```
Not needed
```

_Request Body_

```
Not needed
```

_Response (200)_

```
[
  {
    "id": 1,
    "name": "Ultra Milk",
    "price": 6000,
    "category_id": 1,
    "merchant_id": 1,
    "image": "http://static.bmdstatic.com/pk/product/large/5f87d8756ae79.jpg",
    "stock": 20,
    "created_at": "2021-11-15T14:45:27.93129+08:00",
    "updated_at": "2021-11-15T14:45:27.93129+08:00"
    },
    {
      "id": 2,
      "name": "Milo",
      "price": 4000,
      "category_id": 1,
      "merchant_id": 1,
      "image": "http://static.bmdstatic.com/pk/product/large/5f87d8756ae79.jpg",
      "stock": 80,
      "created_at": "2021-11-15T14:45:27.93129+08:00",
      "updated_at": "2021-11-15T14:45:27.93129+08:00"
    }
]
```


_Response (500)_

```
{  "error" : "Internal Server Error",
  "message": "something went wrong"
}
```

---


<h1 style="color: orange; font-weight: 700;">Merchants</h1>

### POST /merchants
**(create a new merchant data)**

_Request Header_

```
Not needed
```

_Request Body_

```
{   
    "name": "string<merchants's name>",
    "lat": integer<merchants's latitude>,
    "long": integer<merchants's longtitude>,
    "logo": "string<merchants's logo url>"
}
```

_Response (201)_

```
{
    "id": 1,
    "name": "Indomaret Kedoya",
    "lat": "-6.176768",
    "long": "106.764698",
    "logo": "indomaret logo",
    "created_at": "2021-11-15T00:30:20.645569+08:00",
    "updated_at": "2021-11-15T00:30:20.645569+08:00"
}
```

_Response (400)_

```
{ 
  "error"  : "Bad Request"
  "message": "merchant name, lat, long, and logo required"
}
```

_Response (500)_

```
{  "error" : "Internal Server Error",
  "message": "something went wrong"
}
```

---

### GET /merchants/:merchantId
**(get one merchant data by merchantId)**

_Request Header_

```
Not needed
```

_Request Body_

```
Not needed
```

_Response (200)_

```
{
    "id": 1,
    "name": "Indomaret Kedoya",
    "lat": "-6.176768",
    "long": "106.764698",
    "logo": "indomaret logo",
    "created_at": "2021-11-15T00:30:20.645569+08:00",
    "updated_at": "2021-11-15T00:30:20.645569+08:00"
}
```

_Response (400)_

```
{ 
  "error"  : "Bad Request"
  "message": "invalid params id"
}
```

_Response (404)_

```
{ 
  "error"  : "Data Not Found"
  "message": "merchant doesn't exist"
}
```

_Response (500)_

```
{  "error" : "Internal Server Error",
  "message": "something went wrong"
}
```
### PUT /merchants/:merchantId
**(update one merchant data by merchantId )**

_Request Header_

```
Not needed
```

_Request Body_

```
{   
    "name": "string<merchants's name>",
    "lat": integer<merchants's latitude>,
    "long": integer<merchants's longtitude>,
    "logo": "string<merchants's logo url>"
}
```

_Response (200)_

```
{
    "id": 1,
    "name": "Indomaret Kedoya",
    "lat": "-6.176768",
    "long": "106.764698",
    "logo": "indomaret logo",
    "created_at": "2021-11-15T00:30:20.645569+08:00",
    "updated_at": "2021-11-15T00:30:20.645569+08:00"
}
```

_Response (400)_

```
{ 
  "error"  : "Bad Request"
  "message": "merchant name, lat, long, and logo required"
}
```

_Response (500)_

```
{  "error" : "Internal Server Error",
  "message": "something went wrong"
}
```

---

### POST /merchants/get-nearest-merchants
**(get all nearest merchant datas)**

_Request Header_

```
Not needed
```

_Request Body_

```
{   
    "lat": float/double/decimal<latitute>,
    "long": float/double/decimal<longtitude>
}
```

_Response (200)_

```
[
  {
        "id": 1,
        "name": "Indomaret Kedoya",
        "lat": "-6.176768",
        "long": "106.764698",
        "logo": "indomaret logo",
        "created_at": "2021-11-15T00:30:20.645569+08:00",
        "updated_at": "2021-11-15T00:30:20.645569+08:00",
        "distance": 1.6341483293302557,
        "id_2": 1,
        "name_2": "Indomaret Kedoya",
        "lat_2": "-6.176768",
        "long_2": "106.764698",
        "logo_2": "indomaret logo",
        "created_at_2": "2021-11-15T00:30:20.645569+08:00",
        "updated_at_2": "2021-11-15T00:30:20.645569+08:00"
    },
    {
        "id": 5,
        "name": "Indomaret Kedoya",
        "lat": "-6.176768",
        "long": "106.764698",
        "logo": "https://upload.wikimedia.org/wikipedia/id/0/04/Logo_Indomaret.svg",
        "created_at": "2021-11-24T16:20:18.134171+08:00",
        "updated_at": "2021-11-24T16:20:18.134171+08:00",
        "distance": 1.6341483293302557,
        "id_2": 5,
        "name_2": "Indomaret Kedoya",
        "lat_2": "-6.176768",
        "long_2": "106.764698",
        "logo_2": "https://upload.wikimedia.org/wikipedia/id/0/04/Logo_Indomaret.svg",
        "created_at_2": "2021-11-24T16:20:18.134171+08:00",
        "updated_at_2": "2021-11-24T16:20:18.134171+08:00"
    },
]
```

_Response (500)_

```
{  
  "error" : "Internal Server Error",
  "message": "something went wrong"
}
```

---


### PUT /merchants/:merchantId
**(update one merchant data by merchantId )**

_Request Header_

```
Not needed
```

_Request Body_

```
{   
    "name": "string<merchants's name>",
    "lat": integer<merchants's latitude>,
    "long": integer<merchants's longtitude>,
    "logo": "string<merchants's logo url>"
}
```

_Response (200)_

```
{
    "id": 1,
    "name": "Indomaret Kedoya",
    "lat": "-6.176768",
    "long": "106.764698",
    "logo": "indomaret logo",
    "created_at": "2021-11-15T00:30:20.645569+08:00",
    "updated_at": "2021-11-15T00:30:20.645569+08:00"
}
```

_Response (400)_

```
{ 
  "error"  : "Bad Request"
  "message": "merchant name, lat, long, and logo required"
}
```

_Response (500)_

```
{  "error" : "Internal Server Error",
  "message": "something went wrong"
}
```

---

### GET /merchants/
**(get all merchant datas)**

_Request Header_

```
Not needed
```

_Request Body_

```
Not needed
```

_Response (200)_

```
[
    {
        "id": 1,
        "name": "Indomaret Kedoya",
        "lat": "-6.176768",
        "long": "106.764698",
        "logo": "indomaret logo",
        "created_at": "2021-11-15T00:30:20.645569+08:00",
        "updated_at": "2021-11-15T00:30:20.645569+08:00"
    },
    {
        "id": 2,
        "name": "Indomaret Rawajati",
        "lat": "-6.257591",
        "long": "106.851939",
        "logo": "indomaret.logo",
        "created_at": "2021-11-19T14:16:38.709997+08:00",
        "updated_at": "2021-11-19T14:16:38.709997+08:00"
    }
]
```

_Response (500)_

```
{  
  "error" : "Internal Server Error",
  "message": "something went wrong"
}
```

---

<h1 style="color: orange; font-weight: 700;">Carts</h1>

### POST /carts
**(create a new cart data)**

_Request Header_

```
{
  "Authorization": "Bearer <token>"
}
```

_Request Body_

```
{   
  "merchant_id": "integer<cart's merchant_id",
  "product_id": integer<carts's product_id>,
  "qty": integer<carts's quantity>
}
```

_Response (201)_

```
{
    "id": 1,
    "user_id": 1,
    "merchant_id": 2,
    "product_id: 1,
    "qty": 5,
    "total_price": 15000,
    "created_at": "2021-11-15T00:30:20.645569+08:00",
    "updated_at": "2021-11-15T00:30:20.645569+08:00"
}
```

_Response (400)_

```
{ 
  "error"  : "Bad Request"
  "message": "cart's merchant_id, product_id, and qty required"
}
```

_Response (400)_

```
{ 
  "error"  : "Bad Request"
  "message": "insufficient product stock, only 2 left"
}
```

_Response (400)_

```
{ 
  "error"  : "Bad Request"
  "message": "quantity has to be more than or equal to one"
}
```

_Response (401)_

```
{ 
  "error"  : Unauthenticated"
  "message": "invalid token"
}
```

_Response (500)_

```
{  "error" : "Internal Server Error",
  "message": "something went wrong"
}
```

---

### PATCH /carts/:cartId
**(update cart quantity)**

_Request Header_

```
{
  "Authorization": "Bearer <token>"
}
```

_Request Body_

```
{   
  "product_id": integer<carts's product_id>,
  "qty": integer<carts's quantity>
}
```

_Response (200)_

```
{
    "id": 1,
    "user_id": 1,
    "merchant_id": 2,
    "product_id: 1,
    "qty": 5,
    "total_price": 15000,
    "created_at": "2021-11-15T00:30:20.645569+08:00",
    "updated_at": "2021-11-15T00:30:20.645569+08:00"
}
```

_Response (400)_

```
{ 
  "error"  : "Bad Request"
  "message": "cart's product_id, and qty required"
}
```

_Response (400)_

```
{ 
  "error"  : "Bad Request"
  "message": "invalid params id"
}
```

_Response (400)_

```
{ 
  "error"  : "Bad Request"
  "message": "insufficient product stock, only 2 left"
}
```

_Response (400)_

```
{ 
  "error"  : "Bad Request"
  "message": "at least you're purchasing one product"
}
```

_Response (401)_

```
{ 
  "error"  : Unauthenticated"
  "message": "invalid token"
}
```

_Response (403)_

```
{ 
  "error"  : "Unauthorized"
  "message": "you're not allowed to access this data"
}
```

_Response (500)_

```
{  "error" : "Internal Server Error",
  "message": "something went wrong"
}
```

---

---

### Get /carts/my-carts
**(get all carts owned by user)**

_Request Header_

```
{
  "Authorization": "Bearer <token>"
}
```

_Request Body_

```
{   
  "product_id": integer<carts's product_id>,
  "qty": integer<carts's quantity>
}
```

_Response (200)_

```
[
    {
        "id": 3,
        "merchant_name": "Indomaret Kedoya",
        "product_id": 1,
        "product_name": "Ultra Milk",
        "product_price": 6000,
        "product_stock": 20,
        "qty": 5,
        "total_price": 30000,
        "user_id": 1
        "created_at": "2021-11-15T20:43:48.065057+08:00",
        "updated_at": "2021-11-15T20:43:48.065057+08:00",
    }
]

```

_Response (401)_

```
{ 
  "error"  : "Unauthenticated"
  "message": "invalid token"
}
```

_Response (500)_

```
{  "error" : "Internal Server Error",
  "message": "something went wrong"
}
```

---

### DELETE /carts/:cartId
**(delete one cart by cartid)**

_Request Header_

```
{
  "Authorization": "Bearer <token>"
}
```

_Request Body_

```
Not Needed
```

_Response (200)_

```
{
  "message": "cart with id 2 has been sucessfully deleted"
}
```


_Response (400)_

```
{ 
  "error"  : "Bad Request"
  "message": "invalid params id"
}
```



_Response (401)_

```
{ 
  "error"  : Unauthenticated"
  "message": "invalid token"
}
```

_Response (403)_

```
{ 
  "error"  : "Unauthorized"
  "message": "you're not allowed to access this data"
}
```

_Response (500)_

```
{  "error" : "Internal Server Error",
  "message": "something went wrong"
}
```

---

---