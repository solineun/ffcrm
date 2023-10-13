# fulfillment-crm

# REST API

GET /signup -- 200, 500
POST /signup/create_customer --201, 4xx, 500
GET /signin -- 200, 500
POST /signin/auth -- 200, 4xx, 500

GET /customer/:login/orders -- 
GET /customer/:login/orders/:order_id
POST /customer/:login/orders/create
PATCH /customer/:login/orders/:order_id/cancel

GET /customer/:login/products
GET /customer/:login/products/:product_id
POST /customer/:login/products/create
DELETE /customer/:login/products/:product_id/remove

GET /admin/orders
GET /admin/orders/:order_id
PATCH /admin/orders/:order_id/complete
PATCH /admin/orders/:order_id/cancel