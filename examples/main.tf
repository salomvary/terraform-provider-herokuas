terraform {
  required_providers {
    herokuas = {
      version = "0.1.0"
      source  = "github.com/salomvary/herokuas"
    }
  }
}

provider "herokuas" {}

resource "herokuas_trigger" "test_trigger" {
  name           = "Test trigger"
  dyno           = "Free"
  frequency_type = "recurring"
  state          = "active"
  schedule       = "0 1 * * *"
  timezone       = "UTC"
  value          = "echo 'hello world'"
  timeout        = 86400
}