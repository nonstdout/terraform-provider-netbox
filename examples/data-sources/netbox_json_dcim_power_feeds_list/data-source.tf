data "netbox_json_dcim_power_feeds_list" "test" {
  limit = 0
}

output "example" {
  value = jsondecode(data.netbox_json_dcim_power_feeds_list.test.json)
}
