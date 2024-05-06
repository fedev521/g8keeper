# g8keeper

A sample Password Manager (work-in-progress) application on Kubernetes.

Key technologies used:

- Kubernetes
- Helm
- Terraform (to deploy a GKE cluster)
- GitHub Actions
- Prometheus

Table of contents:

- [Components](#components)
- [How to Run Locally](#how-to-run-locally)
  - [Setup](#setup)
  - [Run with Docker Compose](#run-with-docker-compose)
  - [Run with Helm (with Ingress)](#run-with-helm-with-ingress)
  - [Run GitHub Actions](#run-github-actions)
- [How to Monitor with Prometheus and Grafana](#how-to-monitor-with-prometheus-and-grafana)
- [How to Deploy on GKE](#how-to-deploy-on-gke)

## Components

- backend: application server exposing a REST interface to get/list/create
  passwords
- tinksrv: envelope encryption service, it encrypts a DEK without exposing a KEK

## How to Run Locally

### Setup

Create keyset used by tinksrv:

```bash
tinkey create-keyset --key-template=AES256_GCM > keyset.json

# if you want to use docker compose
mv keyset.json src/tinksrv/configs

# if you want to use helm
mv keyset.json helm/g8keeper/secrets/tinksrv
```

Build container images:

```bash
docker compose build
```

Create a local Kubernetes cluster with kind:

```bash
kind create cluster
kind load docker-image tinksrv backend
```

### Run with Docker Compose

```bash
docker compose up
```

### Run with Helm (with Ingress)

```bash
helm upgrade --install ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --namespace ingress-nginx --create-namespace \
  --set controller.service.type=NodePort

helm install helm/g8keeper --generate-name
```

That's it. You can now use curl:

```bash
HTTP_NODE_PORT=$(kubectl get service --namespace ingress-nginx ingress-nginx-controller -o jsonpath="{.spec.ports[0].nodePort}")
HTTPS_NODE_PORT=$(kubectl get service --namespace ingress-nginx ingress-nginx-controller -o jsonpath="{.spec.ports[1].nodePort}")
NODE_IP="$(kubectl get nodes -o jsonpath='{.items[0].status.addresses[?(@.type=="InternalIP")].address}')"

curl \
  --resolve g8keeper.localcluster.me:$HTTP_NODE_PORT:$NODE_IP \
  http://g8keeper.localcluster.me:$HTTP_NODE_PORT/api/v1/passwords

echo -n '{"name":"google", "password":"mY$3kreT"}' \
| curl -s -d @- -X POST http://g8keeper.localcluster.me:$HTTP_NODE_PORT/api/v1/passwords \
    --resolve g8keeper.localcluster.me:$HTTP_NODE_PORT:$NODE_IP

curl \
  --resolve g8keeper.localcluster.me:$HTTP_NODE_PORT:$NODE_IP \
  http://g8keeper.localcluster.me:$HTTP_NODE_PORT/api/v1/passwords

# alternative with port forwarding
kubectl port-forward --namespace=ingress-nginx service/ingress-nginx-controller 8888:80
curl \
  --resolve g8keeper.localcluster.me:8888:127.0.0.1 \
  http://g8keeper.localcluster.me:8888/api/v1/passwords
```

You can render manifests with:

```bash
helm template g8k helm/g8keeper \
  | yq eval 'del(. | select(.kind == "Secret"))' \
  > k8s-manifests/g8keeper-no-secrets.yaml
```

### Run GitHub Actions

```bash
cat > secrets/act.env <<EOT
DOCKERHUB_USERNAME=johndoe
DOCKERHUB_TOKEN=dckr_pat_...
GITHUB_TOKEN=github_pat_...
EOT
act --secret-file secrets/act.env
```

## How to Monitor with Prometheus and Grafana

Follow [these instructions](docs/monitor-with-prom.md).

## How to Deploy on GKE

Follow [these instructions](docs/deploy-on-gke.md).
