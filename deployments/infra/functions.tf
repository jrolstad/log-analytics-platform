resource "oci_functions_application" "primary" {
  compartment_id = oci_identity_compartment.primary.id
  display_name   = "${var.service_name}-${var.environment}"
  subnet_ids     = [oci_core_subnet.primary.id]
  freeform_tags  = { "service" = var.service_name, "environment" = var.environment }
}

data "oci_artifacts_container_images" "file_receiver" {
  compartment_id = oci_functions_application.primary.id
  repository_id  = oci_artifacts_container_repository.primary.id
  version        = var.file_receiver_version
}

resource "oci_functions_function" "file_published_handler" {
  application_id = oci_functions_application.primary.id
  display_name   = "file-published-handler"
  image          = data.oci_artifacts_container_images.file_receiver.container_image_collection[0]
  memory_in_mbs  = 2048
  #config = var.function_config
  timeout_in_seconds = 300
  freeform_tags      = { "service" = var.service_name, "environment" = var.environment }
}