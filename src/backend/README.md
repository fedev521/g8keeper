# Backend Application Server

Run with:

```bash
cd cmd/backend
go run .
```

Build container image and run container with:

```bash
docker build -t backend:latest .
docker run --rm -it backend
```

Call with:

```bash
curl localhost:8080/v1/passwords

echo -n '{"name":"google", "password":"mY$3kreT"}' \
| curl -s -d @- -X POST localhost:8080/v1/passwords

echo -n '{"name":"amazon", "password":"mYP4$$w0rD"}' \
| curl -s -d @- -X POST localhost:8080/v1/passwords

curl localhost:8080/v1/passwords/google
curl localhost:8080/v1/passwords/amazon
```
