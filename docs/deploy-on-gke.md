## Deploy g8keeper on GKE

### How to deploy a GKE cluster (Terraform)

```bash
gcloud auth application-default login

cd terraform/gke

terraform init
terraform apply
# You will be prompted for a few variables, create a terraform.tfvars file based
# on the contents of sample.tfvars. The GCP project is supposed to be already existing.
```

After that, generate kubeconfig with:

```bash
gcloud container clusters get-credentials "$cluster" \
    --location="$location" \
    --project="$project_id"
```

### Deploy the application on GKE (Helm)

Create `gke-values.yaml` based on the contents of `helm/sample-gke-values.yaml`,
then run:

```bash
cd helm

helm upgrade --install ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --namespace ingress-nginx --create-namespace \
# note that controller.service.type will be Load Balancer

helm install ./g8keeper --generate-name -f gke-values.yaml
```

If you can't reach the GKE cluster with a valid domain name, you need to add
something like the following to your `/etc/hosts` file:

```
104.196.143.36 g8keeper.gkecluster.dev`
``
