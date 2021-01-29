resource "bigcommerce_webhook" "example" {
  scope       = "store/customer/*"
  destination = "https://foo.bar/webhook"
  is_active   = true
}