# Tink Demo Server

Run with:

```bash
cd cmd/tinksrv
go run .
```

Build container image and run container with:

```bash
docker build -t tinksrv:latest .
docker run --rm -it tinksrv
```

Call with:

```bash
echo -n '{"plaintext":"encrypt-and-decrypt"}' \
| curl -s -d @- -X POST localhost:8081/v1/encrypt \
| curl -s -d @- -X POST localhost:8081/v1/decrypt
```
