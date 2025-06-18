# bewell-backend-challenge
Backend Golang Coding Test

# Project structure
```bash
bewell-backend-challenge
  |-- cmd
  |-- util
  |-- internal
  |  |-- app
  |  |  |-- port
  |  |  |-- usecase
  |  |-- config
  |  |-- adapter
  |  |  |-- handler
  |  |  |-- router
  |  |-- constant
  |  |-- model
```

## Getting started

### Run local
Requirement
- Go version >= 1.23.8
- Docker

### Setup
Run command:
- go mod vendor
- go mod tidy

ENV in .env file:
```bash
APP_NAME="bewell-backend-challenge"
APPLICATION_NAME=bewell backend challenge
APP_ENV="local"
APP_PORT=8080
VERSION="1.0.0"
BASE_URL=http://localhost
HTTP_CLIENT_TIMEOUT=30
```

Setup database and run service:

    - Create database in docker container

        ```bash
        docker compose up -d
        ```
    
    - Run API
        The server will run on port 8080. (<http://localhost:8080>)

    - Run local

        ```bash
         go run cmd/main.go
        ```

    - Run test

        ```bash
         go test -v ./...
        ```

### Sample API Request/Response
```bash
    - [GET] /ping
        status_code: 200
        response: {
            "message": "pong"
        }

    - [POST] /v1/clean-orders
        status_code: 200
        request: {
            "orders":[
                {
                    "no": 1,
                    "platformProductId": "FG0A-CLEAR-IPHONE16PROMAX",
                    "qty": 2,
                    "unitPrice": 50,
                    "totalPrice": 100
                }
            ]
        }
        response: {
            "result": [
                {
                    "no": 1,
                    "productId": "FG0A-CLEAR-IPHONE16PROMAX",
                    "materialId": "FG0A-CLEAR",
                    "modelId": "IPHONE16PROMAX",
                    "qty": 2,
                    "unitPrice": 50,
                    "totalPrice": 100
                },
                {
                    "no": 2,
                    "productId": "WIPING-CLOTH",
                    "qty": 2,
                    "unitPrice": 0,
                    "totalPrice": 0
                },
                {
                    "no": 3,
                    "productId": "CLEAR-CLEANER",
                    "qty": 2,
                    "unitPrice": 0,
                    "totalPrice": 0
                }
            ]
        }

        status_code: 400
        request: {
            "orders":[
                {
                    "no": 1,
                    "platformProductId": "--FG0A-CLEAR-OPPOA3*2/FG0A-MATTE-OPPOA3*2",
                    "qty": 1,
                    "unitPrice": 160,
                    "totalPrice": 160
                }
            ]
        }
        response: {
            "status_text": "Validation Failed",
            "error": {
                "error_code": "REQUIRED",
                "field": "UnitPrice",
                "message": "Validation failed on field 'UnitPrice', condition: required, actual: 0"
            }
        }
```

### Ex. curl
```bash
    curl --location 'http://localhost:8080/v1/clean-orders' \
    --header 'Content-Type: application/json' \
    --data '{
        "orders":[
            {
                "no": 1,
                "platformProductId": "--FG0A-CLEAR-OPPOA3*2/FG0A-MATTE-OPPOA3*2",
                "qty": 1,
                "unitPrice": 160,
                "totalPrice": 160
            },
            {
                "no": 2,
                "platformProductId": "FG0A-PRIVACY-IPHONE16PROMAX",
                "qty": 1,
                "unitPrice": 50,
                "totalPrice": 50
            }
        ]
    }'
```
