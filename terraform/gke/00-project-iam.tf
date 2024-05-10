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

resource "google_iam_workload_identity_pool" "cicd" {
  project                   = var.project_id
  workload_identity_pool_id = "wi-pool-t1-${var.env}"
}

resource "google_iam_workload_identity_pool_provider" "gha" {
  project                            = var.project_id
  workload_identity_pool_id          = google_iam_workload_identity_pool.cicd.workload_identity_pool_id
  workload_identity_pool_provider_id = "wi-pool-provider-gha-cicd-${var.env}"
  attribute_mapping = {
    "google.subject"             = "assertion.sub"
    "attribute.actor"            = "assertion.actor"
    "attribute.repository"       = "assertion.repository"
    "attribute.repository_owner" = "assertion.repository_owner"
  }
  attribute_condition = "assertion.repository_owner == '${var.github_account}'"
  oidc {
    issuer_uri = "https://token.actions.githubusercontent.com"
  }
}

module "node_sa" {
  source = "github.com/GoogleCloudPlatform/cloud-foundation-fabric//modules/iam-service-account?ref=v30.0.0"

  project_id   = module.project.project_id
  name         = "sa-gkenode-${var.env}"
  display_name = "GKE Nodes"
}

module "cicd_sa" {
  source = "github.com/GoogleCloudPlatform/cloud-foundation-fabric//modules/iam-service-account?ref=v30.0.0"

  project_id   = module.project.project_id
  name         = "sa-gha-cicd-${var.env}"
  display_name = "GitHub Actions"

  iam_bindings_additive = {
    "gha-cicd" = {
      role   = "roles/iam.workloadIdentityUser"
      member = "principalSet://iam.googleapis.com/${google_iam_workload_identity_pool.cicd.name}/attribute.repository/${var.github_account}/${var.github_repo}"
    }
  }
}
