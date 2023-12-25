resource "oci_functions_application" "primary" {
  compartment_id = oci_identity_compartment.primary.id
  display_name   = "${var.service_name}-${var.environment}"
  subnet_ids     = [oci_core_subnet.subnet_1.id]
  freeform_tags  = { "service" = var.service_name, "environment" = var.environment }
}