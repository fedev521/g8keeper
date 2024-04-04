# g8keeper

Password Manager on Kubernetes.

## Architecture

Components:

- backend: application server exposing a REST interface to get/list/create
  passwords
- tinksrv: envelope encryption service, it encrypts a DEK without exposing a KEK

## Setup

```bash
tinkey create-keyset --key-template=AES256_GCM > keyset.json

# if you want to use docker compose
mv keyset.json src/tinksrv/configs

# if you want to use helm
mv keyset.json g8keeper/secrets/tinksrv
```

## How to Run Locally

### Docker Compose

```bash
docker compose up --build
```

### Kind

```bash
kind create cluster
kind load docker-image tinksrv backend

kubectl apply -f k8s/manifests
kubectl port-forward service/backend 8080:8080
```
