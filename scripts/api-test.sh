#!/bin/bash

curl -X POST "http://127.0.0.1:8080/product" \
-H "Content-Type: application/json" \
-d '{
    "name": "Sticker"
}' \
-v

curl -X GET "http://127.0.0.1:8080/product?name=Sticker"