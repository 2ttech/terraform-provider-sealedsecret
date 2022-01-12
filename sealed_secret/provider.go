package sealedsecret

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},
		ResourcesMap: map[string]*schema.Resource{
			"sealed_secret_raw_secret": resourceRawSecret(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
	}
}