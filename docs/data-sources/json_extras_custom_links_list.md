# netbox\_json\_extras\_custom\_links\_list Data Source

Get json output from the extras_custom_links_list Netbox endpoint

## Example Usage

```hcl
data "netbox_json_extras_custom_links_list" "test" {
  limit = 0
}
output "example" {
  value = jsondecode(data.netbox_json_extras_custom_links_list.test.json)
}
```

## Argument Reference

* ``limit`` (Optional). The max number of returned results. If 0 is specified, all records will be returned.

## Attributes Reference

In addition to the above arguments, the following attributes are exported:
* ``json`` - JSON output of the list of objects for this Netbox endpoint.
