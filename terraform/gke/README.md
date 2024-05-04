# GKE Module

Based on [cloud-foundation-fabric](github.com/GoogleCloudPlatform/cloud-foundation-fabric)
Terraform modules.

Creates:

- VPC, subnets, NAT, firewall
- service accounts
  - for GKE nodes
  - for the GitHub Actions pipeline
- IAM permissions
- GKE cluster
