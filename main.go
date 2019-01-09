package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/paultyng/terraform-provider-git/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider})
}
