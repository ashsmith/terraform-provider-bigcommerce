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