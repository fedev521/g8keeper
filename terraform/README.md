# Terraform

## How to deploy a GKE cluster

```bash
gcloud auth application-default login

cd terraform/gke

terraform init
terraform apply
# You will be prompted for a few variables, create a terraform.tfvars file based
# on the contents of sample.tfvars. The GCP is supposed to be already existing.
```

After that, generate kubeconfig with:

```bash
gcloud container clusters get-credentials "$cluster" \
    --location="$zone"
```
