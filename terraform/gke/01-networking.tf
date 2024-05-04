module "vpc" {
  source = "github.com/GoogleCloudPlatform/cloud-foundation-fabric//modules/net-vpc?ref=v30.0.0"

  project_id = module.project.project_id
  name       = var.vpc_name
  vpc_create = var.vpc_create
  subnets = [
    {
      ip_cidr_range = var.mgmt_subnet_cidr_block
      name          = "subnet-mgmt-${var.env}"
      region        = var.region
    },
    {
      ip_cidr_range = var.cluster_network_config.nodes_cidr_block
      name          = "subnet-cluster-${var.env}"
      region        = var.region
      secondary_ip_ranges = {
        pods     = var.cluster_network_config.pods_cidr_block
        services = var.cluster_network_config.services_cidr_block
      }
    }
  ]
}

module "firewall" {
  source     = "github.com/GoogleCloudPlatform/cloud-foundation-fabric//modules/net-vpc-firewall?ref=v30.0.0"
  project_id = module.project.project_id
  network    = module.vpc.name
}

module "nat" {
  source         = "github.com/GoogleCloudPlatform/cloud-foundation-fabric//modules/net-cloudnat?ref=v30.0.0"
  project_id     = module.project.project_id
  region         = var.region
  name           = "nat-${var.env}"
  router_network = module.vpc.name
}

