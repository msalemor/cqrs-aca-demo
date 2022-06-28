# Solution Schemas

[Home](README.md)

## Schemas

### Vendors

Sample Vendor

```json
{
    "vendortId": 1000,
    "name": "product name",
    "phone": "999-999-9999",
    "email": "vendor@email.com",
    "createdDate": "2022-06-22T00:00:00Z"
}
```

### Products

Sample Product

```json
{
    "productId": 1000,
    "name": "product name",
    "vendorId": 1000,
    "price": 10.0,
    "createdDate": "2022-06-22T00:00:00Z"
}
```

### Warehouses

Sample Warehouse

```json
{
    "warehouseId": 1000,
    "name": "product name",
    "city": "Miami",
    "state": "FL"
}
```

### Inventory Transactions

Sample Inventory Transaction

```json
{
    "eventType": "inventory",
    "transactionType": "receiving", //shipping, adjustment
    "warehouseId": 1000,
    "productId": 1000,
    "qty": 10,
    "createdDate": "2022-06-22T00:00:00Z"
}
```