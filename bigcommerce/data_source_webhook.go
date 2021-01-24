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
		ReadContext: dataSourceWebhookRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"client_id": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"store_hash": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"created_at": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"updated_at": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"destination": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_active": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"headers": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
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
	if err := d.Set("headers", webhook.Headers); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
