package herokuas

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HEROKU_ADVANCED_SCHEDULER_API_TOKEN", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"herokuas_trigger": resourceTrigger(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	token := d.Get("token").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	c := NewClient(token)

	return c, diags
}
