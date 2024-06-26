package ipam

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	netboxclient "github.com/nonstdout/go-netbox/v3/netbox/client"
	"github.com/nonstdout/go-netbox/v3/netbox/client/ipam"
	"github.com/nonstdout/go-netbox/v3/netbox/models"
	"github.com/nonstdout/terraform-provider-netbox/v7/netbox/internal/customfield"
	"github.com/nonstdout/terraform-provider-netbox/v7/netbox/internal/tag"
	"github.com/nonstdout/terraform-provider-netbox/v7/netbox/internal/util"
)

func ResourceNetboxIpamService() *schema.Resource {
	return &schema.Resource{
		Description:   "Manage an service (ipam module) within Netbox.",
		CreateContext: resourceNetboxIpamServiceCreate,
		ReadContext:   resourceNetboxIpamServiceRead,
		UpdateContext: resourceNetboxIpamServiceUpdate,
		DeleteContext: resourceNetboxIpamServiceDelete,
		Exists:        resourceNetboxIpamServiceExists,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"content_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The content type of this service (ipam module).",
			},
			"custom_field": &customfield.CustomFieldSchema,
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      nil,
				ValidateFunc: validation.StringLenBetween(1, 200),
				Description:  "The description of this service (ipam module).",
			},
			"device_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				ExactlyOneOf: []string{"device_id", "virtualmachine_id"},
				Description:  "ID of the device linked to this service (ipam module).",
			},
			"ip_addresses_id": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "Array of ID of IP addresses attached to this service (ipam module).",
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 50),
				Description:  "The name for this service (ipam module).",
			},
			"ports": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "Array of ports of this service (ipam module).",
			},
			"protocol": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"tcp", "udp"}, false),
				Description:  "The protocol of this service (ipam module) (tcp or udp).",
			},
			"tag": &tag.TagSchema,
			"virtualmachine_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "ID of the VM linked to this service (ipam module).",
			},
		},
	}
}

func resourceNetboxIpamServiceCreate(ctx context.Context, d *schema.ResourceData,
	m interface{}) diag.Diagnostics {
	client := m.(*netboxclient.NetBoxAPI)

	resourceCustomFields := d.Get("custom_field").(*schema.Set).List()
	customFields := customfield.ConvertCustomFieldsFromTerraformToAPI(nil, resourceCustomFields)
	description := d.Get("description").(string)
	deviceID := int64(d.Get("device_id").(int))
	IPaddressesID := d.Get("ip_addresses_id").([]interface{})
	IPaddressesID64 := []int64{}
	name := d.Get("name").(string)
	ports := d.Get("ports").([]interface{})
	ports64 := []int64{}
	protocol := d.Get("protocol").(string)
	tags := d.Get("tag").(*schema.Set).List()
	virtualmachineID := int64(d.Get("virtualmachine_id").(int))

	for _, id := range IPaddressesID {
		IPaddressesID64 = append(IPaddressesID64, int64(id.(int)))
	}

	for _, p := range ports {
		ports64 = append(ports64, int64(p.(int)))
	}

	newResource := &models.WritableService{
		CustomFields: &customFields,
		Description:  description,
		Ipaddresses:  IPaddressesID64,
		Name:         &name,
		Ports:        ports64,
		Protocol:     &protocol,
		Tags:         tag.ConvertTagsToNestedTags(tags),
	}

	if deviceID != 0 {
		newResource.Device = &deviceID
	}

	if virtualmachineID != 0 {
		newResource.VirtualMachine = &virtualmachineID
	}

	resource := ipam.NewIpamServicesCreateParams().WithData(newResource)

	resourceCreated, err := client.Ipam.IpamServicesCreate(resource, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(resourceCreated.Payload.ID, 10))
	return resourceNetboxIpamServiceRead(ctx, d, m)
}

func resourceNetboxIpamServiceRead(ctx context.Context, d *schema.ResourceData,
	m interface{}) diag.Diagnostics {
	client := m.(*netboxclient.NetBoxAPI)

	resourceID := d.Id()
	params := ipam.NewIpamServicesListParams().WithID(&resourceID)
	resources, err := client.Ipam.IpamServicesList(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	for _, resource := range resources.Payload.Results {
		if strconv.FormatInt(resource.ID, 10) == d.Id() {
			if err = d.Set("content_type", util.ConvertURIContentType(resource.URL)); err != nil {
				return diag.FromErr(err)
			}

			resourceCustomFields := d.Get("custom_field").(*schema.Set).List()
			customFields := customfield.UpdateCustomFieldsFromAPI(resourceCustomFields, resource.CustomFields)

			if err = d.Set("custom_field", customFields); err != nil {
				return diag.FromErr(err)
			}

			var description interface{}
			if resource.Description == "" {
				description = nil
			} else {
				description = resource.Description
			}

			if err = d.Set("description", description); err != nil {
				return diag.FromErr(err)
			}

			if resource.Device == nil {
				if err = d.Set("device_id", nil); err != nil {
					return diag.FromErr(err)
				}
			} else {
				if err = d.Set("device_id", resource.Device.ID); err != nil {
					return diag.FromErr(err)
				}
			}

			IPaddressesObject := resource.Ipaddresses
			IPaddressesInt := []int64{}
			for _, ip := range IPaddressesObject {
				IPaddressesInt = append(IPaddressesInt, ip.ID)
			}
			if err = d.Set("ip_addresses_id", IPaddressesInt); err != nil {
				return diag.FromErr(err)
			}

			if err = d.Set("name", resource.Name); err != nil {
				return diag.FromErr(err)
			}

			ports := resource.Ports
			if err = d.Set("ports", ports); err != nil {
				return diag.FromErr(err)
			}

			if err = d.Set("protocol", resource.Protocol.Value); err != nil {
				return diag.FromErr(err)
			}

			if err = d.Set("tag", tag.ConvertNestedTagsToTags(resource.Tags)); err != nil {
				return diag.FromErr(err)
			}

			if resource.VirtualMachine == nil {
				if err = d.Set("virtualmachine_id", nil); err != nil {
					return diag.FromErr(err)
				}
			} else {
				if err = d.Set("virtualmachine_id", resource.VirtualMachine.ID); err != nil {
					return diag.FromErr(err)
				}
			}

			return nil
		}
	}

	d.SetId("")
	return nil
}

func resourceNetboxIpamServiceUpdate(ctx context.Context, d *schema.ResourceData,
	m interface{}) diag.Diagnostics {
	client := m.(*netboxclient.NetBoxAPI)
	params := &models.WritableService{}

	// Required parameters
	name := d.Get("name").(string)
	params.Name = &name

	ports := d.Get("ports").([]interface{})
	ports64 := []int64{}
	for _, port := range ports {
		ports64 = append(ports64, int64(port.(int)))
	}

	params.Ports = ports64

	protocol := d.Get("protocol").(string)
	params.Protocol = &protocol

	IPaddressesID := d.Get("ip_addresses_id").([]interface{})
	IPaddressesID64 := []int64{}
	for _, id := range IPaddressesID {
		IPaddressesID64 = append(IPaddressesID64, int64(id.(int)))
	}

	params.Ipaddresses = IPaddressesID64

	deviceID := int64(d.Get("device_id").(int))
	if deviceID != 0 {
		params.Device = &deviceID
	}

	vmID := int64(d.Get("virtualmachine_id").(int))
	if vmID != 0 {
		params.VirtualMachine = &vmID
	}

	if d.HasChange("custom_field") {
		stateCustomFields, resourceCustomFields := d.GetChange("custom_field")
		customFields := customfield.ConvertCustomFieldsFromTerraformToAPI(stateCustomFields.(*schema.Set).List(), resourceCustomFields.(*schema.Set).List())
		params.CustomFields = &customFields
	}

	if d.HasChange("description") {
		if description, exist := d.GetOk("description"); exist {
			params.Description = description.(string)
		} else {
			params.Description = " "
		}
	}

	tags := d.Get("tag").(*schema.Set).List()
	params.Tags = tag.ConvertTagsToNestedTags(tags)

	resource := ipam.NewIpamServicesPartialUpdateParams().WithData(
		params)

	resourceID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to convert ID into int64")
	}

	resource.SetID(resourceID)

	_, err = client.Ipam.IpamServicesPartialUpdate(resource, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceNetboxIpamServiceRead(ctx, d, m)
}

func resourceNetboxIpamServiceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*netboxclient.NetBoxAPI)

	resourceExists, err := resourceNetboxIpamServiceExists(d, m)
	if err != nil {
		return diag.FromErr(err)
	}

	if !resourceExists {
		return nil
	}

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.Errorf("Unable to convert ID into int64")
	}

	resource := ipam.NewIpamServicesDeleteParams().WithID(id)
	if _, err := client.Ipam.IpamServicesDelete(resource, nil); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceNetboxIpamServiceExists(d *schema.ResourceData, m interface{}) (b bool,
	e error) {
	client := m.(*netboxclient.NetBoxAPI)
	resourceExist := false

	vlanID := d.Id()
	params := ipam.NewIpamServicesListParams().WithID(&vlanID)
	vlans, err := client.Ipam.IpamServicesList(params, nil)

	if err != nil {
		return resourceExist, err
	}

	for _, vlan := range vlans.Payload.Results {
		if strconv.FormatInt(vlan.ID, 10) == d.Id() {
			resourceExist = true
		}
	}

	return resourceExist, nil
}
