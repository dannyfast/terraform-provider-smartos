package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema:         providerSchema(),
		ResourcesMap:   providerResources(),
		DataSourcesMap: providerDataSources(),
	}
}

// List of supported configuration fields for the provider.
// More info in https://github.com/hashicorp/terraform/blob/v0.6.6/helper/schema/schema.go#L29-L142
func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"host": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "Host address of the SmartOS global zone.",
		},
	}
}

func providerResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"smartos_machine": resourceMachine(),
	}
}

func providerDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"smartos_image": datasourceImage(),
	}
}
