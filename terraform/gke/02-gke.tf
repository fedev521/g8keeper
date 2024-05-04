locals {
  cluster_location = var.zonal ? "${var.region}-c" : var.region
}

module "gke_cluster" {
  source = "github.com/GoogleCloudPlatform/cloud-foundation-fabric//modules/gke-cluster-standard?ref=v30.0.0"

  project_id = module.project.project_id
  name       = var.cluster_name
  location   = local.cluster_location
  vpc_config = {
    network                  = module.vpc.self_link
    subnetwork               = module.vpc.subnet_self_links["${var.region}/subnet-cluster-${var.env}"]
    secondary_range_names    = {}
    master_authorized_ranges = var.cluster_network_config.master_authorized_cidr_blocks
    master_ipv4_cidr_block   = var.cluster_network_config.master_cidr_block
  }

  max_pods_per_node = 32
  private_cluster_config = {
    # enable_private_endpoint: When true, the cluster's private endpoint is used
    # as the cluster endpoint and access through the public endpoint is
    # disabled. When false, either endpoint can be used.
    enable_private_endpoint = !var.enable_public_endpoint
    master_global_access    = false
  }

  deletion_protection = false
}

module "gke_nodepool_default" {
  source = "github.com/GoogleCloudPlatform/cloud-foundation-fabric//modules/gke-nodepool?ref=v30.0.0"

  project_id   = module.project.project_id
  cluster_name = module.gke_cluster.name
  location     = local.cluster_location
  name         = "${module.gke_cluster.name}-nodepool-01"
  service_account = {
    email        = module.node_sa.email
    oauth_scopes = ["https://www.googleapis.com/auth/cloud-platform"]
  }
  node_config = var.default_node_config
}
