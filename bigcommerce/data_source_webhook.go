package bigcommerce

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ashsmith/bigcommerce-api-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWebhook() *schema.Resource {
	return &schema.Resource{
		Description: "Provides information about a webhook ",
		ReadContext: dataSourceWebhookRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"client_id": &schema.Schema{
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"store_hash": &schema.Schema{
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"created_at": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"updated_at": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"destination": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_active": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"header": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceWebhookRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	c := m.(*bigcommerce.Client)
	hookID := d.Get("id").(string)

	fmt.Println(c)

	webhookID, _ := strconv.ParseInt(hookID, 10, 64)
	webhook, whErr := c.Webhooks.Get(webhookID)

	if whErr != nil {
		return diag.FromErr(whErr)
	}

	err := setWebhookData(webhook, d)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(webhook.ID, 10))

	return diags
}

func setWebhookData(webhook bigcommerce.Webhook, d *schema.ResourceData) diag.Diagnostics {
	if err := d.Set("id", strconv.FormatInt(webhook.ID, 10)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("client_id", webhook.ClientID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("store_hash", webhook.StoreHash); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("created_at", int(webhook.CreatedAt)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("updated_at", int(webhook.UpdatedAt)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("scope", webhook.Scope); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("destination", webhook.Destination); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("is_active", webhook.IsActive); err != nil {
		return diag.FromErr(err)
	}

	// Convert webhook.Headers (map[string]string) into compatible slice: []map[string]string [{ key: "", value: ""}]
	headers := make([]interface{}, 0)
	for key, value := range webhook.Headers {
		header := make(map[string]interface{})
		header["key"] = key
		header["value"] = value
		headers = append(headers, header)
	}

	if err := d.Set("header", headers); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
