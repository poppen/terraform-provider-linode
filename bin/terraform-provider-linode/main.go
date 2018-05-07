package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/poppen/terraform-provider-linode"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: linode.Provider,
	})
}
