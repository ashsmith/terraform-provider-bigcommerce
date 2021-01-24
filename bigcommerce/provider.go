package bigcommerce

import (
	"context"
	"net/http"

	"github.com/ashsmith/bigcommerce-api-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"store_hash": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("BIGCOMMERCE_STORE_HASH", nil),
			},
			"client_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("BIGCOMMERCE_CLIENT_ID", nil),
			},
			"access_token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("BIGCOMMERCE_ACCESS_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"bigcommerce_webhook": resourceWebhook(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"bigcommerce_webhook": dataSourceWebhook(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	clientID := d.Get("client_id").(string)
	accessToken := d.Get("access_token").(string)
	storeHash := d.Get("store_hash").(string)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if clientID == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Missing client_id from provider configuration",
			Detail:   "client_id is a required parmeter and must be defined, you can also use BIGCOMMERCE_CLIENT_ID environment variable.",
		})
	}

	if accessToken == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Missing access_token from provider configuration",
			Detail:   "access_token is a required parmeter and must be defined, you can also use BIGCOMMERCE_ACCESS_TOKEN environment variable.",
		})
	}

	if storeHash == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Missing store_hash from provider configuration",
			Detail:   "store_hash is a required parmeter and must be defined, you can also use BIGCOMMERCE_STORE_HASH environment variable.",
		})
	}

	if (clientID == "") || (accessToken == "") || (storeHash == "") {
		return nil, diags
	}

	app := bigcommerce.App{
		ClientID:    clientID,
		StoreHash:   storeHash,
		AccessToken: accessToken,
	}

	httpClient := http.Client{}
	bcClient := app.NewClient(httpClient)

	return bcClient, diags
}
