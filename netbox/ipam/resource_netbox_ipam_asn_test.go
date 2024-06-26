package ipam_test

import (
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/nonstdout/terraform-provider-netbox/v7/netbox/internal/util"
)

const resourceNameNetboxIpamAsn = "netbox_ipam_asn.test"

func TestAccNetboxIpamAsnMinimal(t *testing.T) {
	nameSuffix := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	asn := int64(acctest.RandIntRange(1, 4294967295))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { util.TestAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckNetboxIpamAsnConfig(nameSuffix, false, false, asn),
				Check: resource.ComposeTestCheckFunc(
					util.TestAccResourceExists(resourceNameNetboxIpamAsn),
				),
			},
			{
				ResourceName:      resourceNameNetboxIpamAsn,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxIpamAsnFull(t *testing.T) {
	nameSuffix := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	asn := int64(acctest.RandIntRange(1, 4294967295))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { util.TestAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckNetboxIpamAsnConfig(nameSuffix, true, true, asn),
				Check: resource.ComposeTestCheckFunc(
					util.TestAccResourceExists(resourceNameNetboxIpamAsn),
				),
			},
			{
				ResourceName:      resourceNameNetboxIpamAsn,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxIpamAsnMinimalFullMinimal(t *testing.T) {
	nameSuffix := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	asn := int64(acctest.RandIntRange(1, 4294967295))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { util.TestAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckNetboxIpamAsnConfig(nameSuffix, false, false, asn),
				Check: resource.ComposeTestCheckFunc(
					util.TestAccResourceExists(resourceNameNetboxIpamAsn),
				),
			},
			{
				Config: testAccCheckNetboxIpamAsnConfig(nameSuffix, true, true, asn),
				Check: resource.ComposeTestCheckFunc(
					util.TestAccResourceExists(resourceNameNetboxIpamAsn),
				),
			},
			{
				Config: testAccCheckNetboxIpamAsnConfig(nameSuffix, false, true, asn),
				Check: resource.ComposeTestCheckFunc(
					util.TestAccResourceExists(resourceNameNetboxIpamAsn),
				),
			},
			{
				Config: testAccCheckNetboxIpamAsnConfig(nameSuffix, false, false, asn),
				Check: resource.ComposeTestCheckFunc(
					util.TestAccResourceExists(resourceNameNetboxIpamAsn),
				),
			},
		},
	})
}

func testAccCheckNetboxIpamAsnConfig(nameSuffix string, resourceFull, extraResources bool, asn int64) string {
	template := `
	resource "netbox_ipam_rir" "test" {
		name = "test-{{ .namesuffix }}"
		slug = "test-{{ .namesuffix }}"
	}

	{{ if eq .extraresources "true" }}
	resource "netbox_tenancy_tenant" "test" {
		name = "test-{{ .namesuffix }}"
		slug = "test-{{ .namesuffix }}"
	}

	resource "netbox_extras_tag" "test" {
		name = "test-{{ .namesuffix }}"
		slug = "test-{{ .namesuffix }}"
	}
	{{ end }}

	resource "netbox_ipam_asn" "test" {
		asn         = {{ .asn }}
		rir_id      = netbox_ipam_rir.test.id
		{{ if eq .resourcefull "true" }}
		description = "Test ASN"
		tenant_id   = netbox_tenancy_tenant.test.id

		tag {
			name = netbox_extras_tag.test.name
			slug = netbox_extras_tag.test.slug
		}
		{{ end }}
	}
	`
	data := map[string]string{
		"namesuffix":     nameSuffix,
		"resourcefull":   strconv.FormatBool(resourceFull),
		"extraresources": strconv.FormatBool(extraResources),
		"asn":            strconv.FormatInt(asn, 10),
	}
	return util.RenderTemplate(template, data)
}
