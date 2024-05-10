variable "project_id" {
  description = "GCP project that will contain the Kubernetes cluster."
  type        = string
}

variable "region" {
  description = "GCP region where you will create GCE instances."
  type        = string
  default     = "us-east1"
}

variable "env" {
  description = "Target environment. This concept is up to you, any short string is fine."
  type        = string
  default     = "g8k"
}

variable "github_account" {
  description = "GitHub account name. Used for Workload Identity Federation (GitHub Actions)."
  type        = string
}

variable "github_repo" {
  description = "GitHub repository name. Used for Workload Identity Federation (GitHub Actions)."
  type        = string
}

################################################################################

variable "cluster_network_config" {
  description = "Cluster network configuration."
  type = object({
    nodes_cidr_block              = string
    pods_cidr_block               = string
    services_cidr_block           = string
    master_authorized_cidr_blocks = map(string)
    master_cidr_block             = string
  })
  default = {
    nodes_cidr_block    = "10.0.1.0/24"
    pods_cidr_block     = "172.16.0.0/20"
    services_cidr_block = "192.168.0.0/24"
    master_authorized_cidr_blocks = {
      anyone = "0.0.0.0/0"
    }
    master_cidr_block = "10.0.0.0/28"
  }
}

variable "enable_public_endpoint" {
  type        = bool
  default     = true
  description = <<-EOT
    Whether the control plane's public endpoint is used or not.

    If false, make sure to set appropriate internal IP ranges in `var.cluster_network_config.master_authorized_cidr_blocks`.

    Note that the cluster is private. Basically this flag lets you choose between:

    - External endpoint access disabled
    - External endpoint access enabled, authorized networks enabled

    See https://cloud.google.com/kubernetes-engine/docs/concepts/private-cluster-concept#overview for more information.
  EOT
}

variable "mgmt_subnet_cidr_block" {
  description = "Management subnet IP CIDR range."
  type        = string
  default     = "10.0.2.0/24"
}

variable "vpc_name" {
  description = "VPC name."
  type        = string
  nullable    = false
  default     = "vpc-01"
}

variable "vpc_create" {
  description = "Flag indicating whether the VPC should be created or not."
  type        = bool
  default     = true
}

################################################################################

variable "zonal" {
  description = "Flag indicating whether the cluster should be zonal or regional"
  type        = bool
  default     = false
}

variable "cluster_name" {
  type        = string
  description = "The GKE cluster name"
  default     = "gke-cluster"
}

variable "default_node_config" {
  description = "Node-level configuration of the default nodepool."
  type = object({
    boot_disk_kms_key   = optional(string)
    disk_size_gb        = optional(number)
    disk_type           = optional(string)
    ephemeral_ssd_count = optional(number)
    gcfs                = optional(bool, false)
    guest_accelerator = optional(object({
      count = number
      type  = string
      gpu_driver = optional(object({
        version                    = string
        partition_size             = optional(string)
        max_shared_clients_per_gpu = optional(number)
      }))
    }))
    local_nvme_ssd_count = optional(number)
    gvnic                = optional(bool, false)
    image_type           = optional(string)
    kubelet_config = optional(object({
      cpu_manager_policy   = string
      cpu_cfs_quota        = optional(bool)
      cpu_cfs_quota_period = optional(string)
      pod_pids_limit       = optional(number)
    }))
    linux_node_config = optional(object({
      sysctls     = optional(map(string))
      cgroup_mode = optional(string)
    }))
    local_ssd_count       = optional(number)
    machine_type          = optional(string)
    metadata              = optional(map(string))
    min_cpu_platform      = optional(string)
    preemptible           = optional(bool, false)
    sandbox_config_gvisor = optional(bool)
    shielded_instance_config = optional(object({
      enable_integrity_monitoring = optional(bool)
      enable_secure_boot          = optional(bool)
    }))
    spot                          = optional(bool, false)
    workload_metadata_config_mode = optional(string)
  })
  default = {
    disk_type    = "pd-balanced"
    machine_type = "n2-standard-2"
    disk_size_gb = 50
  }
}
