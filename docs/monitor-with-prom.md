# Monitor with Prometheus and Grafana

## Setup

Just see `helm/prometheus/install.sh`.

## Access Grafana

If you are using Ingress rules, visit
http://grafana.localcluster.me:$HTTP_NODE_PORT/. You have to update `/etc/hosts`
accordingly.

If you want to use port forwarding, run the following command, then visit
http://localhost:4040/:

```bash
kubectl port-forward services/prometheus-grafana 4040:80 -n prometheus
```

The default credentials are the following:

```
admin
prom-operator
```
