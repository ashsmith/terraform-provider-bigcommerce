---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bigcommerce_webhook Data Source - terraform-provider-bigcommerce"
subcategory: ""
description: |-
  Provides information about a webhook
---

# Data Source `bigcommerce_webhook`

Provides information about a webhook

## Example Usage

```terraform
data "bigcommerce_webhook" "example" {
  id = "123456"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) The ID of this resource.

### Read-only

- **client_id** (String, Sensitive)
- **created_at** (Number)
- **destination** (String)
- **header** (Block Set) (see [below for nested schema](#nestedblock--header))
- **is_active** (Boolean)
- **scope** (String)
- **store_hash** (String, Sensitive)
- **updated_at** (Number)

<a id="nestedblock--header"></a>
### Nested Schema for `header`

Read-only:

- **key** (String)
- **value** (String)


