# terraform-provider-bigcommerce

Terraform Provider for BigCommerce

- [Bigcommerce Provider Documentation on Terraform](https://registry.terraform.io/providers/ashsmith/bigcommerce/latest)

## Resources:

- `bigcommerce_webhook` - The ability to create webhooks for your

## Data Sources:

- `bigcommerce_webhook` - Ability to read an existing webhook.

## Build & install locally:

```
make install
```

## Generating docs

```
go generate
```

## Example usage:

```terraform
terraform {
  required_providers {
    bigcommerce = {
      source = "ashsmith/bigcommerce"
      version = "0.1.0"
    }
  }
}

provider "bigcommerce" {
  store_hash   = "your-hash"
  client_id    = "your-client-id"
  access_token = "your-access-token"
}

resource "bigcommerce_webhook" "example" {
  scope       = "store/customer/*"
  destination = "https://foo.bar/webhook"
  is_active   = true

  header {
    key   = "X-My-Header"
    value = "myheadervalue"
  }

  header {
    key   = "X-My-Other-Header"
    value = "myheadervalue"
  }
}
```

## Environment variables

The bigcommerce provider also supports environment variables in place of the provider configuration options. This makes it easier to apply your terraform code across multipe environment without the use of tfvars.

```
export BIGCOMMERCE_STORE_HASH=your-hash
export BIGCOMMERCE_CLIENT_ID=your-client-id
export BIGCOMMERCE_ACCESS_TOKEN=your-access-token
terraform apply
```
