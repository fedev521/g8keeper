module "project" {
  source = "github.com/GoogleCloudPlatform/cloud-foundation-fabric//modules/project?ref=v30.0.0"

  project_create = false
  name           = var.project_id
  services = [
    "artifactregistry.googleapis.com",
    "container.googleapis.com",
    "compute.googleapis.com"
  ]
  iam = {
    "roles/container.nodeServiceAccount" = [module.node_sa.iam_email]
    "roles/container.developer"          = [module.cicd_sa.iam_email]
  }

  skip_delete = true
}

module "node_sa" {
  source = "github.com/GoogleCloudPlatform/cloud-foundation-fabric//modules/iam-service-account?ref=v30.0.0"

  project_id = module.project.project_id
  name       = "sa-gkenode-${var.env}"
}

module "cicd_sa" {
  source = "github.com/GoogleCloudPlatform/cloud-foundation-fabric//modules/iam-service-account?ref=v30.0.0"

  project_id = module.project.project_id
  name       = "sa-gha-cicd-${var.env}"
}
