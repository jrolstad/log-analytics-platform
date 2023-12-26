resource "oci_logging_log_group" "primary" {
  compartment_id = oci_identity_compartment.primary.id
  display_name   = "${var.service_name}-${var.environment}"
  freeform_tags  = { "service" = var.service_name, "environment" = var.environment }
}

resource "oci_logging_log" "connector" {
  display_name = "connectorhub-runlog"
  log_group_id = oci_logging_log_group.primary.id
  log_type     = "SERVICE"

  configuration {
    compartment_id = oci_identity_compartment.primary.id
    source {
      category    = "runlog"
      resource    = oci_sch_service_connector.file_published_handler.id
      service     = "sch.serviceconnector"
      source_type = "OCISERVICE"
    }
  }
  freeform_tags      = { "service" = var.service_name, "environment" = var.environment }
  is_enabled         = true
  retention_duration = 30
}