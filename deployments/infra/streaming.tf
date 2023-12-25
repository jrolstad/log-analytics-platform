resource "oci_streaming_stream_pool" "primary" {
  compartment_id = oci_identity_compartment.primary.id
  name           = "${var.service_name}-${var.environment}"
  freeform_tags  = { "service" = var.service_name, "environment" = var.environment }
}

resource "oci_streaming_stream" "file_published" {
  name               = "file_published"
  partitions         = var.default_stream_partion_count
  compartment_id     = oci_identity_compartment.primary.id
  freeform_tags      = { "service" = var.service_name, "environment" = var.environment }
  retention_in_hours = 24
  stream_pool_id     = oci_streaming_stream_pool.primary.id
}