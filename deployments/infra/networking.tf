resource "oci_core_virtual_network" "primary" {
  compartment_id = oci_identity_compartment.primary.id
  cidr_block     = "10.0.0.0/16"
  display_name   = "${var.service_name}-${var.environment}"
  freeform_tags  = { "service" = var.service_name, "environment" = var.environment }
}

resource "oci_core_subnet" "primary" {
  compartment_id = oci_identity_compartment.primary.id
  vcn_id         = oci_core_virtual_network.primary.id
  cidr_block     = "10.0.0.0/24"
  freeform_tags  = { "service" = var.service_name, "environment" = var.environment }
}

data "oci_core_services" "all" {
}

resource "oci_core_service_gateway" "primary" {
  compartment_id = oci_identity_compartment.primary.id
  vcn_id         = oci_core_virtual_network.primary.id
  dynamic "services" {
    for_each = data.oci_core_services.all.services
    content {
      service_id = each.id
    }
  }
  freeform_tags = { "service" = var.service_name, "environment" = var.environment }
}