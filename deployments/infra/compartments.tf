resource "oci_identity_compartment" "primary" {
  compartment_id = var.root_compartment_id
  description    = "Resources for the ${var.service_name} service, supporting the ${var.environment} environment"
  name           = "${var.service_name}-${var.environment}"
  enable_delete  = true

  freeform_tags = { "service" = var.service_name, "environment" = var.environment }
}