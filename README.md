# User

Handles authorithation, jwt token creation and check.

## How tu run

``` SH
docker build -t user_api .

docker run --rm -d --name user_api \
  -e APP_HOST=0.0.0.0 \
  -e APP_PORT=8080 \
  -e DB_CONN="postgres://user:pass@localhost:5432/user?sslmode=disable" \
  -p 8000:8080 user_api
```
