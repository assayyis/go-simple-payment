# How To Run
- Go to project path
- Run `go get .` to get the necessary packages
- Run `go run .` to start the project
- Project running at `localhost:8080`

# API Endpoints
## Login
Endpoint: `api/auth/login`

Request Body
```json
{
  "username": "customer1",
  "password": "password"
}
```

## Payment
Endpoint: `api/payment`

Request Body
```json
{
  "CustomerId": 1,
  "MerchantId": 1,
  "Amount": 5000
}
```

## Logout
Endpoint: `api/auth/logout/:id`
