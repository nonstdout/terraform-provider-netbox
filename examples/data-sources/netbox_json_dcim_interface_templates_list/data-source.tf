data "netbox_json_dcim_interface_templates_list" "test" {
  limit = 0
}

output "example" {
  value = jsondecode(data.netbox_json_dcim_interface_templates_list.test.json)
}
