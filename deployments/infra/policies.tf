resource "oci_identity_dynamic_group" "FunctionsServiceDynamicGroup" {
  name           = "FunctionsServiceDynamicGroup-${var.service_name}-${var.environment}"
  description    = "Dynamic Group for the functions in the ${var.service_name}-${var.environment} compartment"
  compartment_id = var.tenant_id
  matching_rule  = "ALL {resource.type = 'fnfunc', resource.compartment.id = '${oci_identity_compartment.primary.id}'}"

  provisioner "local-exec" {
    command = "sleep 5"
  }
}

resource "oci_identity_policy" "FunctionsServiceDynamicGroupPolicy" {
  depends_on     = [oci_identity_dynamic_group.FunctionsServiceDynamicGroup]
  name           = "FunctionsServiceDynamicGroupPolicy-${var.service_name}-${var.environment}"
  description    = "Allows functions to manage all resources in the ${var.service_name}-${var.environment} compartment"
  compartment_id = oci_identity_compartment.primary.id
  statements     = ["allow dynamic-group ${oci_identity_dynamic_group.FunctionsServiceDynamicGroup.name} to manage all-resources in compartment id ${oci_identity_compartment.primary.id}"]

  provisioner "local-exec" {
    command = "sleep 5"
  }
}