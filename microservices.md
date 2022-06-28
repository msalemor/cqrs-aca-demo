# Solution Microservices

[Home](README.md)

## Command Stack (Go)

### Vendors

- POST /vendor
- PUT /vendor/:id
- DELETE /vendor/:id

### Products

- POST /product
- PUT /product/:id
- DELETE /product/:id

### Warehouse

- POST /Warehouse
- PUT /Warehouse/:id
- DELETE /Warehouse/:id

## Event Source stack (Python)

### Inventory

- PUT /inventory/ [single or array]

## Snapshot stack

### Vendors (timer) (Go)

- Vendors -> Vendors snapshot

### Products (timer) (Go)

- Products -> Products snapshot

### Inventory (event) (Python)

- Write to store
- Rebuild Inventory snapshot

## Query stack (Go)

### Vendors

GET /vendor
GET /vendor/:id

### Products

GET /product
GET /product/:id

### Warehouse

GET /Warehouse
GET /Warehouse/:id

### Inventory

GET /inventory
GET /inventory/:id

## Frontend (React)