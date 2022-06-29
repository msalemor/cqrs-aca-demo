# Solution Microservices

[Home](README.md)

## Command Stack (Go)

### Vendors (Go/Fiber)

- POST /vendor
- PUT /vendor/:id
- DELETE /vendor/:id

### Products (Go/Fiber)

- POST /product
- PUT /product/:id
- DELETE /product/:id

### Warehouse (Go/Fiber)

- POST /Warehouse
- PUT /Warehouse/:id
- DELETE /Warehouse/:id

## Event Source stack (Python)

### Inventory (Python/ServiceBus)

- PUT /inventory/ [single or array]

## Snapshot stack

### Vendors (timer) 

- Vendors -> Vendors snapshot

### Products (timer) (Go)

- Products -> Products snapshot

### Inventory (event) (Python)

- Write to store
- Rebuild Inventory snapshot

## Query stack (Go)

### Vendors (Go/Fiber)

GET /vendor
GET /vendor/:id

### Products (Go/Fiber)

GET /product
GET /product/:id

### Warehouse (Go/Fiber)

GET /Warehouse
GET /Warehouse/:id

### Inventory (Go/Fiber)

GET /inventory
GET /inventory/:id

## Frontend (React)
