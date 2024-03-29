# g8keeper

Password Manager on Kubernetes.

## Architecture

Components:

- backend: application server exposing a REST interface to get/list/create
  passwords
- tinksrv: envelope encryption service, it encrypts a DEK without exposing a KEK

## Setup

```bash
src/tinksrv
tinkey create-keyset --key-template=AES256_GCM > configs/keyset.json
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
