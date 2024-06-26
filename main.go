package main

import (
	"flag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/nonstdout/terraform-provider-netbox/v7/netbox"
)

// Run "go generate" to format example terraform files and generate the docs for the registry/website

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		Debug:        debug,
		ProviderAddr: "registry.terraform.io/nonstdout/netbox",
		ProviderFunc: netbox.Provider,
	}

	plugin.Serve(opts)
}
