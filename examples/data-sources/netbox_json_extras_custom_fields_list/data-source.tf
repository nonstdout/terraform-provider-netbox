data "netbox_json_extras_custom_fields_list" "test" {
  limit = 0
}

output "example" {
  value = jsondecode(data.netbox_json_extras_custom_fields_list.test.json)
}
