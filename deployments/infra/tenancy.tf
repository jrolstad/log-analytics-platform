data "oci_identity_tenancy" "primary" {
  tenancy_id = var.tenant_id
}