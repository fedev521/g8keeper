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

echo -n '{"name":"Google", "password":"mY$3kreT"}' \
| curl -s -d @- -X POST localhost:8080/v1/passwords
```
