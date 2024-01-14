resource "oci_core_virtual_network" "primary" {
  compartment_id = oci_identity_compartment.primary.id
  cidr_block     = "10.0.0.0/16"
  display_name   = "${var.service_name}-${var.environment}"
  freeform_tags  = { "service" = var.service_name, "environment" = var.environment }
}

resource "oci_core_internet_gateway" "primary" {
  compartment_id = oci_identity_compartment.primary.id
  display_name   = "${var.service_name}-${var.environment}"
  vcn_id         = oci_core_virtual_network.primary.id
}

resource "oci_core_route_table" "primary" {
  compartment_id = oci_identity_compartment.primary.id
  display_name   = "${var.service_name}-${var.environment}"
  vcn_id         = oci_core_virtual_network.primary.id
  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = oci_core_internet_gateway.primary.id
  }
}

resource "oci_core_dhcp_options" "primary" {
  compartment_id = oci_identity_compartment.primary.id
  display_name   = "${var.service_name}-${var.environment}"
  vcn_id         = oci_core_virtual_network.primary.id

  // required
  options {
    type        = "DomainNameServer"
    server_type = "VcnLocalPlusInternet"
  }

}

resource "oci_core_subnet" "primary" {
  compartment_id    = oci_identity_compartment.primary.id
  display_name      = "${var.service_name}-${var.environment}"
  vcn_id            = oci_core_virtual_network.primary.id
  cidr_block        = "10.0.0.0/24"
  dns_label         = "fnsub"
  route_table_id    = oci_core_route_table.primary.id
  dhcp_options_id   = oci_core_dhcp_options.primary.id
  security_list_ids = [oci_core_virtual_network.primary.default_security_list_id]

  freeform_tags = { "service" = var.service_name, "environment" = var.environment }
}

data "oci_core_services" "all" {
}

resource "oci_core_service_gateway" "primary" {
  compartment_id = oci_identity_compartment.primary.id
  vcn_id         = oci_core_virtual_network.primary.id
  services {
    service_id = data.oci_core_services.all.services[0].id
  }
  freeform_tags = { "service" = var.service_name, "environment" = var.environment }
}