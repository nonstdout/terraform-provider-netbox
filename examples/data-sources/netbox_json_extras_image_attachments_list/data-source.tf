data "netbox_json_extras_image_attachments_list" "test" {
  limit = 0
}

output "example" {
  value = jsondecode(data.netbox_json_extras_image_attachments_list.test.json)
}
