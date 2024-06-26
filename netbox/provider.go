package netbox

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/nonstdout/go-netbox/v3/netbox/client"
	"github.com/nonstdout/terraform-provider-netbox/v7/netbox/dcim"
	"github.com/nonstdout/terraform-provider-netbox/v7/netbox/extras"
	"github.com/nonstdout/terraform-provider-netbox/v7/netbox/ipam"
	"github.com/nonstdout/terraform-provider-netbox/v7/netbox/json"
	"github.com/nonstdout/terraform-provider-netbox/v7/netbox/tenancy"
	"github.com/nonstdout/terraform-provider-netbox/v7/netbox/virtualization"
)

const authHeaderName = "Authorization"
const authHeaderFormat = "Token %v"

// Provider exports the actual provider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETBOX_URL", "127.0.0.1:8000"),
				Description: "URL and port to reach netbox application (127.0.0.1:8000 by default).",
			},
			"basepath": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETBOX_BASEPATH", client.DefaultBasePath),
				Description: "URL base path to the netbox API (/api by default).",
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETBOX_TOKEN", ""),
				Description: "Token used for API operations (empty by default).",
			},
			"scheme": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETBOX_SCHEME", "https"),
				Description: "Scheme used to reach netbox application (https by default).",
			},
			"insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETBOX_INSECURE", false),
				Description: "Skip TLS certificate validation (false by default).",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"netbox_json_circuits_circuit_terminations_list":   json.DataNetboxJSONCircuitsCircuitTerminationsList(),
			"netbox_json_circuits_circuit_types_list":          json.DataNetboxJSONCircuitsCircuitTypesList(),
			"netbox_json_circuits_circuits_list":               json.DataNetboxJSONCircuitsCircuitsList(),
			"netbox_json_circuits_provider_networks_list":      json.DataNetboxJSONCircuitsProviderNetworksList(),
			"netbox_json_circuits_providers_list":              json.DataNetboxJSONCircuitsProvidersList(),
			"netbox_json_dcim_cables_list":                     json.DataNetboxJSONDcimCablesList(),
			"netbox_json_dcim_console_server_ports_list":       json.DataNetboxJSONDcimConsoleServerPortsList(),
			"netbox_json_dcim_device_roles_list":               json.DataNetboxJSONDcimDeviceRolesList(),
			"netbox_json_dcim_device_types_list":               json.DataNetboxJSONDcimDeviceTypesList(),
			"netbox_json_dcim_devices_list":                    json.DataNetboxJSONDcimDevicesList(),
			"netbox_json_dcim_interface_templates_list":        json.DataNetboxJSONDcimInterfaceTemplatesList(),
			"netbox_json_dcim_inventory_item_templates_list":   json.DataNetboxJSONDcimInventoryItemTemplatesList(),
			"netbox_json_dcim_locations_list":                  json.DataNetboxJSONDcimLocationsList(),
			"netbox_json_dcim_manufacturers_list":              json.DataNetboxJSONDcimManufacturersList(),
			"netbox_json_dcim_module_bay_templates_list":       json.DataNetboxJSONDcimModuleBayTemplatesList(),
			"netbox_json_dcim_module_types_list":               json.DataNetboxJSONDcimModuleTypesList(),
			"netbox_json_dcim_modules_list":                    json.DataNetboxJSONDcimModulesList(),
			"netbox_json_dcim_platforms_list":                  json.DataNetboxJSONDcimPlatformsList(),
			"netbox_json_dcim_power_feeds_list":                json.DataNetboxJSONDcimPowerFeedsList(),
			"netbox_json_dcim_power_ports_list":                json.DataNetboxJSONDcimPowerPortsList(),
			"netbox_json_dcim_racks_list":                      json.DataNetboxJSONDcimRacksList(),
			"netbox_json_dcim_rear_port_templates_list":        json.DataNetboxJSONDcimRearPortTemplatesList(),
			"netbox_json_dcim_rear_ports_list":                 json.DataNetboxJSONDcimRearPortsList(),
			"netbox_json_dcim_sites_list":                      json.DataNetboxJSONDcimSitesList(),
			"netbox_json_dcim_virtual_chassis_list":            json.DataNetboxJSONDcimVirtualChassisList(),
			"netbox_json_dcim_virtual_device_contexts_list":    json.DataNetboxJSONDcimVirtualDeviceContextsList(),
			"netbox_json_extras_config_contexts_list":          json.DataNetboxJSONExtrasConfigContextsList(),
			"netbox_json_extras_content_types_list":            json.DataNetboxJSONExtrasContentTypesList(),
			"netbox_json_extras_custom_fields_list":            json.DataNetboxJSONExtrasCustomFieldsList(),
			"netbox_json_extras_custom_links_list":             json.DataNetboxJSONExtrasCustomLinksList(),
			"netbox_json_extras_export_templates_list":         json.DataNetboxJSONExtrasExportTemplatesList(),
			"netbox_json_extras_image_attachments_list":        json.DataNetboxJSONExtrasImageAttachmentsList(),
			"netbox_json_extras_job_results_list":              json.DataNetboxJSONExtrasJobResultsList(),
			"netbox_json_extras_journal_entries_list":          json.DataNetboxJSONExtrasJournalEntriesList(),
			"netbox_json_extras_object_changes_list":           json.DataNetboxJSONExtrasObjectChangesList(),
			"netbox_json_extras_saved_filters_list":            json.DataNetboxJSONExtrasSavedFiltersList(),
			"netbox_json_extras_tags_list":                     json.DataNetboxJSONExtrasTagsList(),
			"netbox_json_extras_webhooks_list":                 json.DataNetboxJSONExtrasWebhooksList(),
			"netbox_json_ipam_aggregates_list":                 json.DataNetboxJSONIpamAggregatesList(),
			"netbox_json_ipam_asns_list":                       json.DataNetboxJSONIpamAsnsList(),
			"netbox_json_ipam_fhrp_group_assignments_list":     json.DataNetboxJSONIpamFhrpGroupAssignmentsList(),
			"netbox_json_ipam_fhrp_groups_list":                json.DataNetboxJSONIpamFhrpGroupsList(),
			"netbox_json_ipam_ip_addresses_list":               json.DataNetboxJSONIpamIPAddressesList(),
			"netbox_json_ipam_ip_ranges_list":                  json.DataNetboxJSONIpamIPRangesList(),
			"netbox_json_ipam_l2vpn_terminations_list":         json.DataNetboxJSONIpamL2vpnTerminationsList(),
			"netbox_json_ipam_l2vpns_list":                     json.DataNetboxJSONIpamL2vpnsList(),
			"netbox_json_ipam_prefixes_list":                   json.DataNetboxJSONIpamPrefixesList(),
			"netbox_json_ipam_rirs_list":                       json.DataNetboxJSONIpamRirsList(),
			"netbox_json_ipam_roles_list":                      json.DataNetboxJSONIpamRolesList(),
			"netbox_json_ipam_route_targets_list":              json.DataNetboxJSONIpamRouteTargetsList(),
			"netbox_json_ipam_service_templates_list":          json.DataNetboxJSONIpamServiceTemplatesList(),
			"netbox_json_ipam_services_list":                   json.DataNetboxJSONIpamServicesList(),
			"netbox_json_ipam_vlan_groups_list":                json.DataNetboxJSONIpamVlanGroupsList(),
			"netbox_json_ipam_vlans_list":                      json.DataNetboxJSONIpamVlansList(),
			"netbox_json_ipam_vrfs_list":                       json.DataNetboxJSONIpamVrfsList(),
			"netbox_json_tenancy_contact_assignments_list":     json.DataNetboxJSONTenancyContactAssignmentsList(),
			"netbox_json_tenancy_contact_groups_list":          json.DataNetboxJSONTenancyContactGroupsList(),
			"netbox_json_tenancy_contact_roles_list":           json.DataNetboxJSONTenancyContactRolesList(),
			"netbox_json_tenancy_contacts_list":                json.DataNetboxJSONTenancyContactsList(),
			"netbox_json_tenancy_tenant_groups_list":           json.DataNetboxJSONTenancyTenantGroupsList(),
			"netbox_json_tenancy_tenants_list":                 json.DataNetboxJSONTenancyTenantsList(),
			"netbox_json_users_groups_list":                    json.DataNetboxJSONUsersGroupsList(),
			"netbox_json_users_permissions_list":               json.DataNetboxJSONUsersPermissionsList(),
			"netbox_json_users_tokens_list":                    json.DataNetboxJSONUsersTokensList(),
			"netbox_json_users_users_list":                     json.DataNetboxJSONUsersUsersList(),
			"netbox_json_virtualization_cluster_groups_list":   json.DataNetboxJSONVirtualizationClusterGroupsList(),
			"netbox_json_virtualization_cluster_types_list":    json.DataNetboxJSONVirtualizationClusterTypesList(),
			"netbox_json_virtualization_clusters_list":         json.DataNetboxJSONVirtualizationClustersList(),
			"netbox_json_virtualization_interfaces_list":       json.DataNetboxJSONVirtualizationInterfacesList(),
			"netbox_json_virtualization_virtual_machines_list": json.DataNetboxJSONVirtualizationVirtualMachinesList(),
			"netbox_json_wireless_wireless_lan_groups_list":    json.DataNetboxJSONWirelessWirelessLanGroupsList(),
			"netbox_json_wireless_wireless_lans_list":          json.DataNetboxJSONWirelessWirelessLansList(),
			"netbox_json_wireless_wireless_links_list":         json.DataNetboxJSONWirelessWirelessLinksList(),
			"netbox_dcim_device_role":                          dcim.DataNetboxDcimDeviceRole(),
			"netbox_dcim_location":                             dcim.DataNetboxDcimLocation(),
			"netbox_dcim_manufacturer":                         dcim.DataNetboxDcimManufacturer(),
			"netbox_dcim_platform":                             dcim.DataNetboxDcimPlatform(),
			"netbox_dcim_rack":                                 dcim.DataNetboxDcimRack(),
			"netbox_dcim_rack_role":                            dcim.DataNetboxDcimRackRole(),
			"netbox_dcim_region":                               dcim.DataNetboxDcimRegion(),
			"netbox_dcim_site":                                 dcim.DataNetboxDcimSite(),
			"netbox_ipam_aggregate":                            ipam.DataNetboxIpamAggregate(),
			"netbox_ipam_asn":                                  ipam.DataNetboxIpamAsn(),
			"netbox_ipam_ip_addresses":                         ipam.DataNetboxIpamIPAddresses(),
			"netbox_ipam_ip_range":                             ipam.DataNetboxIpamIPRange(),
			"netbox_ipam_prefix":                               ipam.DataNetboxIpamPrefix(),
			"netbox_ipam_rir":                                  ipam.DataNetboxIpamRir(),
			"netbox_ipam_role":                                 ipam.DataNetboxIpamRole(),
			"netbox_ipam_route_targets":                        ipam.DataNetboxIpamRouteTargets(),
			"netbox_ipam_service":                              ipam.DataNetboxIpamService(),
			"netbox_ipam_vlan":                                 ipam.DataNetboxIpamVlan(),
			"netbox_ipam_vlan_group":                           ipam.DataNetboxIpamVlanGroup(),
			"netbox_ipam_vrf":                                  ipam.DataNetboxIpamVrf(),
			"netbox_tenancy_contact":                           tenancy.DataNetboxTenancyContact(),
			"netbox_tenancy_contact_group":                     tenancy.DataNetboxTenancyContactGroup(),
			"netbox_tenancy_contact_role":                      tenancy.DataNetboxTenancyContactRole(),
			"netbox_tenancy_tenant":                            tenancy.DataNetboxTenancyTenant(),
			"netbox_tenancy_tenant_group":                      tenancy.DataNetboxTenancyTenantGroup(),
			"netbox_virtualization_cluster":                    virtualization.DataNetboxVirtualizationCluster(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"netbox_dcim_device_role":             dcim.ResourceNetboxDcimDeviceRole(),
			"netbox_dcim_location":                dcim.ResourceNetboxDcimLocation(),
			"netbox_dcim_manufacturer":            dcim.ResourceNetboxDcimManufacturer(),
			"netbox_dcim_platform":                dcim.ResourceNetboxDcimPlatform(),
			"netbox_dcim_rack":                    dcim.ResourceNetboxDcimRack(),
			"netbox_dcim_rack_role":               dcim.ResourceNetboxDcimRackRole(),
			"netbox_dcim_region":                  dcim.ResourceNetboxDcimRegion(),
			"netbox_dcim_site":                    dcim.ResourceNetboxDcimSite(),
			"netbox_extras_custom_field":          extras.ResourceNetboxExtrasCustomField(),
			"netbox_extras_tag":                   extras.ResourceNetboxExtrasTag(),
			"netbox_ipam_aggregate":               ipam.ResourceNetboxIpamAggregate(),
			"netbox_ipam_asn":                     ipam.ResourceNetboxIpamASN(),
			"netbox_ipam_ip_addresses":            ipam.ResourceNetboxIpamIPAddresses(),
			"netbox_ipam_ip_range":                ipam.ResourceNetboxIpamIPRange(),
			"netbox_ipam_prefix":                  ipam.ResourceNetboxIpamPrefix(),
			"netbox_ipam_rir":                     ipam.ResourceNetboxIpamRIR(),
			"netbox_ipam_route_targets":           ipam.ResourceNetboxIpamRouteTargets(),
			"netbox_ipam_service":                 ipam.ResourceNetboxIpamService(),
			"netbox_ipam_vlan":                    ipam.ResourceNetboxIpamVlan(),
			"netbox_ipam_vlan_group":              ipam.ResourceNetboxIpamVlanGroup(),
			"netbox_ipam_vrf":                     ipam.ResourceNetboxIpamVrf(),
			"netbox_tenancy_contact":              tenancy.ResourceNetboxTenancyContact(),
			"netbox_tenancy_contact_assignment":   tenancy.ResourceNetboxTenancyContactAssignment(),
			"netbox_tenancy_contact_group":        tenancy.ResourceNetboxTenancyContactGroup(),
			"netbox_tenancy_contact_role":         tenancy.ResourceNetboxTenancyContactRole(),
			"netbox_tenancy_tenant":               tenancy.ResourceNetboxTenancyTenant(),
			"netbox_tenancy_tenant_group":         tenancy.ResourceNetboxTenancyTenantGroup(),
			"netbox_virtualization_cluster":       virtualization.ResourceNetboxVirtualizationCluster(),
			"netbox_virtualization_cluster_group": virtualization.ResourceNetboxVirtualizationClusterGroup(),
			"netbox_virtualization_cluster_type":  virtualization.ResourceNetboxVirtualizationClusterType(),
			"netbox_virtualization_interface":     virtualization.ResourceNetboxVirtualizationInterface(),
			"netbox_virtualization_vm":            virtualization.ResourceNetboxVirtualizationVM(),
			"netbox_virtualization_vm_primary_ip": virtualization.ResourceNetboxVirtualizationVMPrimaryIP(),
		},
		ConfigureContextFunc: configureProvider,
	}
}

func configureProvider(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	url := d.Get("url").(string)
	basepath := d.Get("basepath").(string)
	token := d.Get("token").(string)
	scheme := d.Get("scheme").(string)
	insecure := d.Get("insecure").(bool)

	defaultScheme := []string{scheme}

	t := runtimeclient.New(url, basepath, defaultScheme)
	if insecure {
		t.Transport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: insecure} // #nosec G402
	}
	t.DefaultAuthentication = runtimeclient.APIKeyAuth(authHeaderName, "header",
		fmt.Sprintf(authHeaderFormat, token))

	return client.New(t, strfmt.Default), nil
}
