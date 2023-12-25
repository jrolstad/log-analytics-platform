terraform {
  required_providers {
    oci = {
      source  = "oracle/oci"
      version = "5.10.0"
    }
  }
  backend "http" {
    update_method = "PUT"
  }
}

provider "oci" {
  tenancy_ocid     = var.tenant_id
  user_ocid        = var.user_id
  private_key_path = var.auth_keyfile
  fingerprint      = var.auth_fingerprint
  region           = var.region
}