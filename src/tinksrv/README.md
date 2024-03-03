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
curl localhost:8080/

echo -n '{"plaintext":"encrypt-and-decrypt"}' \
| curl -s -d @- -X POST localhost:8080/v1/encrypt \
| curl -s -d @- -X POST localhost:8080/v1/decrypt
```

## Setup

```bash
tinkey create-keyset --key-template=AES256_GCM > configs/keyset.json
```
