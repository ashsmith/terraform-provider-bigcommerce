resource "bigcommerce_webook" "example" {
  scope       = "store/customer/*"
  destination = "https://foo.bar/webhook"
  is_active   = true
}