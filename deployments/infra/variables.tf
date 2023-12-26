variable "tenant_id" {
  description = "Tenant the resources are associated with"
  type        = string
}

variable "root_compartment_id" {
  description = "Root Compartment Id for the tenancy"
  type        = string
}

variable "user_id" {
  description = "User the IaC process is running as"
  type        = string
}

variable "region" {
  description = "OCI region to deploy resources to"
  type        = string
}

variable "auth_fingerprint" {
  description = "Fingerprint of the certificate the user authenticates as"
  type        = string
}

variable "auth_keyfile" {
  description = "Path to the keyfile the user authenticates with"
  type        = string
}

variable "service_name" {
  description = "Name of the service"
  type        = string
  default     = "log-analyzer"
}

variable "environment" {
  description = "What type of environment (dev,tst,prd)"
  type        = string
}

variable "default_stream_partion_count" {
  description = "Default number of stream partitions"
  type        = number
  default     = 1
}

variable "file_receiver_version" {
  description = "Version of the file receiver to use"
  type        = string
  default     = "latest"
}