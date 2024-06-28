terraform {
  required_providers {
    netbox= {
        source = "registry.terraform.io/nonstdout/netbox"
    }
  }
}
provider netbox {
  # Environment variable NETBOX_URL
  url = "netboxuat.isam.corp"

  # Environment variable NETBOX_TOKEN
  token = "f172fc9993498dfd22132c5c40666b7683c216d7"

  # Environment variable NETBOX_SCHEME
  scheme = "https"

  # Environment variable NETBOX_INSECURE
  insecure = "false"
}

resource "netbox_virtualization_vm" "vm" {
  cluster_id = 1
  name       = "dinf51114uk2"
  comments   = "Created by Terraform"
#   role_id    = var.netbox_dcim_device_role
  vcpus      = 2
  disk       = 100
  memory     = 8192
  site_id = 5
  status = "active"
}

resource "netbox_virtualization_vm_primary_ip" "name" {
  virtualmachine_id = 1293
  primary_ip4_id = "5462"
    

  
  
  #primary_ip6_id = netbox_ipam_ip_addresses.ip6_test.id
}

# data "netbox_ipam_ip_addresses" "ipstuff" {
#     address = "10.244.30.53/24"
  
# }

# resource "netbox_ipam_ip_addresses" "ip_addr" {
#     prefix = "10.244.30.53/24"
  
# }

# output "ip" {
#     description = "an ip address"
#   value = data.netbox_ipam_ip_addresses.ipstuff
# }