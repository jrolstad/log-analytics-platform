resource "oci_sch_service_connector" "file_published_handler" {
  #Required
  compartment_id = oci_identity_compartment.primary.id
  display_name   = "${var.service_name}-${var.environment}-filepublished-handler"
  source {
    kind = "streaming"
    cursor {
      kind = "TRIM_HORIZON"
    }
    stream_id = oci_streaming_stream.file_published.id
  }
  target {
    kind        = "functions"
    function_id = oci_functions_function.file_published_handler.id
  }

  freeform_tags = { "service" = var.service_name, "environment" = var.environment }
}