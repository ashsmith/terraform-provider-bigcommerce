package bigcommerce

import (
	"context"
	"strconv"

	"github.com/ashsmith/bigcommerce-api-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWebhook() *schema.Resource {
	return &schema.Resource{
		Description:   "Provides a BigCommerce Webhook resource.",
		CreateContext: resourceWebhookCreate,
		ReadContext:   resourceWebhookRead,
		UpdateContext: resourceWebhookUpdate,
		DeleteContext: resourceWebhookDelete,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
				Required: true,
			},
			"destination": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"is_active": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"header": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceWebhookCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*bigcommerce.Client)
	var diags diag.Diagnostics

	scope := d.Get("scope").(string)
	destination := d.Get("destination").(string)
	isActive := d.Get("is_active").(bool)

	webhook := bigcommerce.Webhook{
		Scope:       scope,
		Destination: destination,
		IsActive:    isActive,
	}

	webhook.Headers = formatHeaders(d)

	result, err := client.Webhooks.Create(webhook)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(int(result.ID)))

	return diags
}

func resourceWebhookRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*bigcommerce.Client)
	var diags diag.Diagnostics

	webhookID, _ := strconv.ParseInt(d.Id(), 10, 64)

	webhook, whErr := client.Webhooks.Get(webhookID)
	if whErr != nil {
		return diag.FromErr(whErr)
	}

	err := setWebhookData(webhook, d)
	if err != nil {
		return err
	}

	return diags
}

func resourceWebhookUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*bigcommerce.Client)

	webhookID, _ := strconv.ParseInt(d.Id(), 10, 64)

	if d.HasChange("scope") || d.HasChange("destination") || d.HasChange("is_active") || d.HasChange("header") {
		webhook := bigcommerce.Webhook{
			ID:          webhookID,
			Scope:       d.Get("scope").(string),
			Destination: d.Get("destination").(string),
			IsActive:    d.Get("is_active").(bool),
		}

		webhook.Headers = formatHeaders(d)

		_, err := client.Webhooks.Update(webhook)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return resourceWebhookRead(ctx, d, m)
}

func resourceWebhookDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*bigcommerce.Client)
	var diags diag.Diagnostics

	webhookID, _ := strconv.ParseInt(d.Id(), 10, 64)

	err := client.Webhooks.Delete(webhookID)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func formatHeaders(d *schema.ResourceData) map[string]string {
	wbHeaders := make(map[string]string)
	headers := d.Get("header").(*schema.Set).List()
	for _, item := range headers {
		header := item.(map[string]interface{})
		wbHeaders[header["key"].(string)] = header["value"].(string)
	}
	return wbHeaders
}
