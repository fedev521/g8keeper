#!/bin/bash

helm upgrade --install prometheus kube-prometheus-stack \
  --repo https://prometheus-community.github.io/helm-charts \
  --namespace prometheus --create-namespace \
  -f prometheus/values.yaml
